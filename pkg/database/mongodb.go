package database

import (
	"context"
	"errors"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	bson2 "gopkg.in/mgo.v2/bson"
	"log"
	"reflect"
	"strconv"
	"time"
)

// Mb3MongoDB represents a mongodb connection and should implement [MB3Database]
type Mb3MongoDB struct {
	user     string          // user name for database
	pwd      string          // password for database
	host     string          // host name of database connection
	dbname   string          // database name
	port     uint16          // port number
	connStr  string          // connection string for a mongodb database
	database *mongo.Database // handle to the database
	dirty    bool            // if true, the database connection was changed and the database will reconnect.
}

var metadata *MB3MetaData

func (db *Mb3MongoDB) GetMetaData() (*MB3MetaData, error) {
	if metadata != nil {
		return metadata, nil
	}
	var query = `
	    [
       	{ "$match": {"deprecated": null}},
        {
            "$facet": {
                "IsomerCount": [
                    {
                        "$project": {
                            "ikey": {
                                "$filter": {
                                    "input": "$compound.link",
                                    "as": "link",
                                    "cond": { "$eq": ["$$link.key", "INCHIKEY"]
                                    }
                                }
                    }}},
                    {
                        "$unwind": "$ikey"
                    },
                    {
                        "$group": {
                            "_id": "$ikey.value",
                        }
                    },
                    {
                        "$count": "count"
                    },
                ],
                "CompoundCount": [
                    {
                        "$project": {
                            "ikey": {
                                "$filter": {
                                    "input": "$compound.link",
                                    "as": "link",
                                    "cond": { "$eq": ["$$link.key", "INCHIKEY"]
                                    }
                                }
                            }}},
                    {
                        "$unwind": "$ikey"
                    },
                    {
                        "$group": {
                            "_id": {"$substr": ["$ikey.value",0,14]},
                        }
                    },
                    {
                        "$count": "count"
                    },
                ],
                "SpectraCount": [
                    {
                        "$count": "count"
                    },
                ]
            }
        },
        {
            "$project": {
                "IsomerCount": {"$first": "$IsomerCount.count"},
                "CompoundCount": {"$first": "$CompoundCount.count"},
                "SpectraCount": {"$first": "$SpectraCount.count"}
            }
        }
    ]`
	var bdoc interface{}
	if err := bson2.UnmarshalJSON([]byte(query), &bdoc); err != nil {
		return nil, err
	}
	cur, err := db.database.Collection(mbCollection).Aggregate(context.Background(), bdoc)
	if err != nil {
		return nil, err
	}
	var dataMap []bson.M
	if err = cur.All(context.Background(), &dataMap); err != nil {
		return nil, err
	}
	md := db.database.Collection(mbMetaCollection).FindOne(context.Background(), bson.D{})
	var mdMap bson.M
	md.Decode(&mdMap)
	time, err := time.Parse(time.RFC3339, mdMap["timestamp"].(string))
	result := MB3MetaData{
		StoredMetadata: MB3StoredMetaData{
			Version:   mdMap["version"].(string),
			TimeStamp: time.UTC(),
			GitCommit: mdMap["commit"].(string),
		},
		SpectraCount:  int(dataMap[0]["SpectraCount"].(int32)),
		CompoundCount: int(dataMap[0]["CompoundCount"].(int32)),
		IsomerCount:   int(dataMap[0]["IsomerCount"].(int32)),
	}
	metadata = &result
	return &result, nil
}

func (db *Mb3MongoDB) GetUniqueValues(filters Filters) (MB3Values, error) {
	match := getQuery(filters)
	var query = `
[
    { "$facet": {
		"Contributor": [
    {
        "$group": {
            "_id": "$contributor",
            "Count": {"$sum": 1}
        },
    },
    {
        "$project": {
            "_id": 0,
            "Val": "$_id",
            "Count": 1
        }
    },
    {
        "$sort": {
            "Val": 1
        }
    }
],
        "InstrumentType": [
            {
                "$group": {
                    "_id": "$acquisition.instrumenttype",
                    "Count": {"$sum": 1}
                },
            },
            {
                "$project": {
                    "_id": 0,
                    "Val": "$_id",
                    "Count": 1
                }
            },
			{
				"$sort": {
					"Val": 1
				}
            }
        ],
        "IonMode": [
            {
                "$unwind": "$acquisition.massspectrometry"
            },
            {
                "$redact": {
                    "$cond": {
                        "if": {"$eq": ["$acquisition.massspectrometry.key", "ION_MODE"]},
                        "then": "$$KEEP",
                        "else": "$$PRUNE"
                    }
                }
            },
            {
                "$group": {
                    "_id": "$acquisition.massspectrometry.value",
                    "Count": {"$sum": 1}
                }
            },
            {
                "$project": {
                    "_id": 0,
                    "Val": "$_id",
                    "Count": 1
                }
            },
			{
				"$sort": {
					"Val": 1
				}
            }
        ],
        "MSType": [
            {
                "$unwind": "$acquisition.massspectrometry"
            },
            {
                "$redact": {
                    "$cond": {
                        "if": {"$eq": ["$acquisition.massspectrometry.key", "MS_TYPE"]},
                        "then": "$$KEEP",
                        "else": "$$PRUNE"
                    }
                }
            },
            {
                "$group": {
                    "_id": "$acquisition.massspectrometry.value",
                    "Count": {"$sum": 1}
                }
            },
            {
                "$project": {
                    "_id": 0,
                    "Val": "$_id",
                    "Count": 1
                }
            },
			{
				"$sort": {
					"Val": 1
				}
            }
        ]    ,

        "Intensity": [
            { "$unwind": "$peak.peak.intensity"},
    {
        "$group": {
            "_id": null,
            "Max": {"$max": "$peak.peak.intensity"},
            "Min": {"$min": "$peak.peak.intensity"}

        }
    },
            {
                "$project": {"_id":0}
            }
        ],
            "Peak": [
                { "$unwind": "$peak.peak.mz"},
                {
                    "$group": {
                        "_id": null,
                        "Max": {"$max": "$peak.peak.mz"},
                        "Min": {"$min": "$peak.peak.mz"}

                    }
                },
                            {
                "$project": {"_id":0}
            }
            ],
            "Mass": [
                {
                    "$group": {
                        "_id": null,
                        "Min": {"$min": "$compound.mass"},
                        "Max": {"$max": "$compound.mass"}

                    }
                },
                {
                    "$project": {"_id":0}
                }
            ]
        }
    }
    ]`
	var result MB3Values
	var bdoc interface{}
	if err := bson2.UnmarshalJSON([]byte(query), &bdoc); err != nil {
		return result, err
	}
	if len(match) > 0 {
		bdoc = append([]interface{}{bson.D{{"$match", match}}}, bdoc.([]interface{})...)
	}
	cur, err := db.database.Collection(mbCollection).Aggregate(context.Background(), bdoc, &options.AggregateOptions{Collation: &options.Collation{Locale: "en"}})
	if err != nil {
		return result, err
	}
	var temp = make([]bson.D, 0)
	if err = cur.All(context.Background(), &temp); err != nil {
		return result, err
	}
	doc, err := bson.Marshal(temp[0])
	if err != nil {
		return result, err
	}
	type MB3ValuesTemp struct {
		InstrumentType []MBCountValues
		MSType         []MBCountValues
		IonMode        []MBCountValues
		Contributor    []MBCountValues
		Intensity      []MBMinMaxValues
		Mass           []MBMinMaxValues
		Peak           []MBMinMaxValues
	}
	var tempV MB3ValuesTemp
	err = bson.Unmarshal(doc, &tempV)
	result = MB3Values{
		Contributor:    tempV.Contributor,
		InstrumentType: tempV.InstrumentType,
		MSType:         tempV.MSType,
		IonMode:        tempV.IonMode,
		Intensity:      tempV.Intensity[0],
		Mass:           tempV.Mass[0],
		Peak:           tempV.Peak[0],
	}
	return result, err
}

// collection names
const (
	mbCollection        = "massbank"
	mbMetaCollection    = "mb_metadata"
	MongoConnectTimeout = "20s"
)

// NewMongoDB creates a mongodb database handle implementing [MB3Database] from the configuration.
// It does test the connection or connect to the database. This should be done by [Connect()].
//
// Returns an error if the connection data is not valid.
func NewMongoDB(config DBConfig) (*Mb3MongoDB, error) {
	if config.Database != MongoDB {
		return nil, errors.New("database type must be Postgres")
	}
	if len(config.DbName) < 1 ||
		len(config.DbConnStr) < 1 &&
			(len(config.DbHost) < 1 ||
				config.DbPort == 0 ||
				config.DbPort > uint(^uint16(0))) {
		return nil, errors.New("database name, host and port not in config, but DbConnStr is also empty")
	}
	if len(config.DbConnStr) > 0 {
		log.Println("Using connection string for database connection, ignoring other values")
	}
	return &Mb3MongoDB{
		user:     config.DbUser,
		pwd:      config.DbPwd,
		host:     config.DbHost,
		dbname:   config.DbName,
		port:     uint16(config.DbPort),
		connStr:  config.DbConnStr,
		database: nil,
		dirty:    true,
	}, nil
}

// Connect see [MB3Database.Connect]
func (db *Mb3MongoDB) Connect() error {
	timeout, _ := time.ParseDuration(MongoConnectTimeout)
	ctx := context.Background()
	if db.dirty && db.database != nil {
		err := db.database.Client().Disconnect(ctx)
		if err != nil {
			return err
		}
		db.database = nil
	}
	if db.database == nil {
		var err error
		var dbClient *mongo.Client
		if len(db.connStr) > 0 {
			if dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI(db.connStr).SetConnectTimeout(timeout)); err != nil {
				return err
			}
		} else {

			clOptions := options.Client().SetAuth(options.Credential{
				AuthMechanism:           "SCRAM-SHA-256",
				AuthMechanismProperties: nil,
				AuthSource:              "admin",
				Username:                db.user,
				Password:                db.pwd,
				PasswordSet:             true,
			}).SetAppName("MassBank3API").SetHosts([]string{db.host + ":" + strconv.FormatInt(int64(db.port), 10)}).SetConnectTimeout(timeout).SetTLSConfig(nil)
			if dbClient, err = mongo.Connect(ctx, clOptions); err != nil {
				return err
			}
		}
		mongoDb := dbClient.Database(db.dbname)
		if err != nil {
			return err
		}
		db.database = mongoDb
		db.dirty = false
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := db.database.Client().Ping(ctx, nil); err != nil {
		db.database = nil
		return err
	}
	return db.init()
}

func (db *Mb3MongoDB) Ping() error {
	timeout, _ := time.ParseDuration(MongoConnectTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if db.database == nil {
		return errors.New("Database not connected")
	}
	client := db.database.Client()
	if client == nil {
		return errors.New("Database not connected")
	}
	return db.database.Client().Ping(ctx, nil)
}

// Disconnect see [MB3Database.Disconnect]
func (db *Mb3MongoDB) Disconnect() error {
	if db.database == nil {
		return errors.New("database not connected")
	}
	err := db.database.Client().Disconnect(context.Background())
	db.database = nil
	return err
}

// Count see [MB3Database.Count]
func (db *Mb3MongoDB) Count() (int64, error) {
	return db.database.Collection(mbCollection).EstimatedDocumentCount(context.Background())
}

// DropAllRecords see [MB3Database.DropAllRecords]
func (db *Mb3MongoDB) DropAllRecords() error {
	if err := db.database.Drop(context.Background()); err != nil {
		return err
	}
	return db.init()
}

// GetRecord see [MB3Database.GetRecord]
func (db *Mb3MongoDB) GetRecord(s *string) (*massbank.MassBank2, error) {
	var bsonResult bson.M
	err := db.database.Collection(mbCollection).FindOne(context.Background(), bson.D{{"accession", s}}).Decode(&bsonResult)
	if err != nil {
		return nil, err
	}
	mb, err := unmarshal2Massbank(err, &bsonResult)
	if err != nil {
		return nil, err
	}
	return mb, err
}

func (db *Mb3MongoDB) GetSmiles(accession *string) (*string, error) {
	rec, err := db.GetRecord(accession)
	if err != nil {
		return nil, err
	}
	return rec.Compound.Smiles, nil
}

// GetRecords see [MB3Database.GetRecords]
func (db *Mb3MongoDB) GetRecords(
	filters Filters,
) (*SearchResult, error) {
	if db.database == nil {
		return nil, errors.New("database not ready")
	}
	var err error
	var cur *mongo.Cursor
	var groupStage = bson.D{{"$group", bson.D{
		{"_id", "$compound.inchi"},
		{"formula", bson.D{{"$addToSet", "$compound.formula"}}},
		{"mass", bson.D{{"$addToSet", "$compound.mass"}}},
		{"names", bson.D{{"$push", "$compound.names"}}},
		{"smiles", bson.D{{"$addToSet", "$compound.smiles"}}},
		{"spectra", bson.D{{"$push", bson.D{{"title", "$recordtitle"}, {"id", "$accession"}}}}},
		{"specCount", bson.D{{"$count", bson.D{}}}},
	}}}
	if filters.Limit <= 0 {
		filters.Limit = DefaultValues.Limit
	}
	if filters.MassEpsilon == nil {
		filters.MassEpsilon = &DefaultValues.MassEpsilon
	}
	if filters.IntensityCutoff == nil {
		filters.IntensityCutoff = &DefaultValues.IntensityCutoff
	}
	if filters.Offset < 0 {
		filters.Offset = DefaultValues.Offset
	}
	var peakDifferenceStages []bson.D
	if filters.PeakDifferences != nil && len(*filters.PeakDifferences) > 0 {
		peakDifferenceStages = db.GetPeakDifferenceResult(filters.PeakDifferences, *filters.MassEpsilon, *filters.IntensityCutoff, filters.Limit, filters.Offset)
	}
	query := getQuery(filters)
	matchstage := bson.D{{"$match", query}}
	sortstage := bson.D{{"$sort", bson.D{{"spectra.0.title", 1}}}}
	projectstage := bson.D{{"$project", bson.D{
		{"names", 1},
		{"formula", 1},
		{"mass", 1},
		{"spectra", 1},
		{"smiles", 1},
		{"inchi", "$_id"},
		{"_id", 0},
	},
	}}
	skipstage := bson.D{{"$skip", filters.Offset}}
	limitstage := bson.D{{"$limit", filters.Limit}}
	facets := bson.D{{"$facet", bson.D{
		{"ResultCount", mongo.Pipeline{bson.D{{"$count", "count"}}}},
		{"SpecCount", mongo.Pipeline{bson.D{{"$group", bson.D{{"_id", nil}, {"count", bson.D{{"$sum", "$specCount"}}}}}}}},
		{"records", mongo.Pipeline{projectstage, sortstage, skipstage, limitstage}},
	}}}
	pipeline := mongo.Pipeline{matchstage}
	if peakDifferenceStages != nil {
		for _, d := range peakDifferenceStages {
			pipeline = append(pipeline, d)
		}
	}
	pipeline = append(pipeline, groupStage, facets)
	cur, err = db.database.Collection(mbCollection).Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	var bsonResult []bson.M
	if err := cur.All(context.Background(), &bsonResult); err != nil {
		return nil, err
	}
	var mbResult = SearchResult{
		SpectraCount: 0,
		ResultCount:  0,
		Data:         map[string]SearchResultData{},
	}
	if len(bsonResult) > 0 && bsonResult[0]["records"] != nil {
		for _, val := range bsonResult[0]["records"].(bson.A) {
			valM := val.(bson.M)
			inchi := valM["inchi"].(string)
			mb, err := unmarshal2SearchResult(valM)
			if err != nil {
				return nil, err
			}
			mbResult.Data[inchi] = *mb
		}
	}
	if len(bsonResult) > 0 && bsonResult[0]["ResultCount"] != nil && len(bsonResult[0]["ResultCount"].(bson.A)) > 0 {
		mbResult.ResultCount = int(bsonResult[0]["ResultCount"].(bson.A)[0].(bson.M)["count"].(int32))
	} else {
		mbResult.ResultCount = 0
	}
	if len(bsonResult) > 0 && bsonResult[0]["SpecCount"] != nil && len(bsonResult[0]["SpecCount"].(bson.A)) > 0 {
		mbResult.SpectraCount = int(bsonResult[0]["SpecCount"].(bson.A)[0].(bson.M)["count"].(int32))
	} else {
		mbResult.SpectraCount = 0
	}
	return &mbResult, nil
}

func getQuery(filters Filters) bson.D {
	result := bson.D{}
	eps := 0.3
	if !filters.IncludeDeprecated {
		result = append(result, bson.E{"deprecated", bson.M{"$eq": nil}})
	}
	if filters.MassEpsilon != nil {
		eps = *filters.MassEpsilon
	}
	if filters.InstrumentType != nil {
		arr := bson.A{}
		for _, it := range *filters.InstrumentType {
			arr = append(arr, it)
		}
		result = append(result,
			bson.E{"acquisition.instrumenttype", bson.M{"$in": arr}})

	}
	if filters.MsType != nil {
		arr := bson.A{}
		for _, it := range *filters.MsType {
			arr = append(arr, bson.D{bson.E{"key", "MS_TYPE"}, bson.E{"value", it}})
		}
		result = append(result,
			bson.E{"acquisition.massspectrometry", bson.M{"$in": arr}})
	}
	if filters.IonMode != massbank.ANY {
		im := bson.E{"acquisition.massspectrometry", bson.D{bson.E{"key", "ION_MODE"}, bson.E{"value", filters.IonMode}}}
		result = append(result, im)
	}
	if filters.Splash != "" {
		result = append(result, bson.E{"peak.splash", filters.Splash})
	}
	if filters.Mass != nil {
		result = append(result, bson.E{"compound.mass", bson.M{"$gt": *filters.Mass - eps, "$lt": *filters.Mass + eps}})
	}
	if filters.CompoundName != "" {
		result = append(result, bson.E{"compound.names", bson.M{"$regex": filters.CompoundName, "$options": "is"}})
	}
	if filters.Formula != "" {
		result = append(result, bson.E{"compound.formula", bson.M{"$regex": filters.Formula, "$options": "is"}})
	}
	if filters.InchiKey != "" {
		im := bson.E{"compound.link", bson.D{bson.E{"key", "INCHIKEY"}, bson.E{"value", filters.InchiKey}}}
		result = append(result, im)
	}
	if filters.PeakDifferences != nil {
		pd := bson.E{"peak.peak.mz.1", bson.M{"$exists": true}}
		result = append(result, pd)
	}
	if filters.Contributor != nil {
		arr := bson.A{}
		for _, co := range *filters.Contributor {
			arr = append(arr, co)
		}
		result = append(result,
			bson.E{"contributor", bson.M{"$in": arr}})
	}
	if filters.Peaks != nil {
		for _, p := range *filters.Peaks {
			pval := bson.M{"$elemMatch": bson.M{"$gte": p - eps, "$lte": p + eps}}
			result = append(result, bson.E{"peak.peak.mz", pval})
		}
	}
	return result
}

func (db *Mb3MongoDB) GetPeakDifferenceResult(search *[]float64, tolerance float64, intensityOffset int64, limit int64, skip int64) []bson.D {
	query := []bson.D{
		{{"$addFields", bson.D{{
			"peak.peak.diff", bson.D{{
				"$function", bson.D{{
					"body", "function(mz,search,tol) {" +
						"let result = {};" +
						"for (let i=0; i < mz.length; i++) { for (let j=i+1; j < mz.length; j++) {" +
						"for (let p=0; p < search.length; p++) {" +
						"if (Math.abs(Math.abs(mz[i].mz - mz[j].mz)-search[p]) <= tol) {" +
						"return true;}}}}" +
						"return false;}"},
					{"args", bson.A{
						bson.D{{"$filter", bson.D{
							{"input", bson.D{
								{"$map", bson.D{
									{"input", "$peak.peak.rel"},
									{"as", "rel"},
									{"in", bson.D{
										{"rel", "$$rel"},
										{"mz",
											bson.D{{"$arrayElemAt",
												bson.A{"$peak.peak.mz", bson.D{{"$indexOfArray", bson.A{"$peak.peak.rel", "$$rel"}}}}}}}}}}}}},
							{"as", "relmz"},
							{"cond", bson.D{{"$gte", bson.A{"$$relmz.rel", intensityOffset}}}}}}},
						search,
						tolerance}},
					{"lang", "js"}}}}}}}},
		{{"$match", bson.D{{"peak.peak.diff", true}}}},
	}
	return query
}

// UpdateMetadata see [MB3Database.UpdateMetadata]
func (db *Mb3MongoDB) UpdateMetadata(meta *massbank.MbMetaData) (string, error) {
	if db.database == nil {
		return "", errors.New("database not ready")
	}
	var res bson.M
	err := db.database.Collection(mbMetaCollection).FindOne(context.Background(), bson.D{{"commit", meta.Commit}, {"timestamp", meta.Timestamp}}, options.FindOne().SetProjection(bson.D{{"_id", 1}})).Decode(&res)
	if err != nil {
		iRes, err := db.database.Collection(mbMetaCollection).InsertOne(context.Background(), meta)
		if err != nil {
			return "", err
		}
		return iRes.InsertedID.(primitive.ObjectID).Hex(), nil

	}
	return res["_id"].(primitive.ObjectID).Hex(), nil
}

// AddRecord see [MB3Database.AddRecord]
func (db *Mb3MongoDB) AddRecord(record *massbank.MassBank2, metadataId string) error {
	if db.database == nil {
		return errors.New("database not ready")
	}
	record.Metadata.VersionRef = metadataId
	_, err := db.database.Collection(mbCollection).InsertOne(context.Background(), *record)
	if err != nil {
		return err
	}
	return nil
}

// AddRecords see [MB3Database.AddRecords]
func (db *Mb3MongoDB) AddRecords(records []*massbank.MassBank2, metadataId string) error {
	if db.database == nil {
		return errors.New("database not ready")
	}
	var recordsI = make([]interface{}, len(records))
	for i, record := range records {
		record.Metadata.VersionRef = metadataId
		recordsI[i] = record
	}
	_, err := db.database.Collection(mbCollection).InsertMany(context.Background(), recordsI)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRecord see [MB3Database.UpdateRecord]
func (db *Mb3MongoDB) UpdateRecord(
	record *massbank.MassBank2,
	metadataId string,
	upsert bool,
) (uint64, uint64, error) {
	if db.database == nil {
		return 0, 0, errors.New("database not ready")
	}
	record.Metadata.VersionRef = metadataId
	opt := options.ReplaceOptions{}
	res, err := db.database.Collection(mbCollection).ReplaceOne(context.Background(), bson.D{{"accession", *record.Accession}}, record, opt.SetUpsert(upsert))
	if err != nil {
		return 0, 0, err
	}
	return uint64(res.ModifiedCount), uint64(res.UpsertedCount), nil
}

// UpdateRecords see [MB3Database.UpdateRecords]
func (db *Mb3MongoDB) UpdateRecords(
	records []*massbank.MassBank2,
	metadataId string,
	upsert bool,
) (uint64, uint64, error) {
	if db.database == nil {
		return 0, 0, errors.New("database not ready")
	}
	var upserted uint64 = 0
	var modified uint64 = 0
	for _, record := range records {
		u, m, err := db.UpdateRecord(record, metadataId, upsert)
		if err != nil {
			return 0, 0, err
		}
		upserted += u
		modified += m
	}
	return upserted, modified, nil
}

func (db *Mb3MongoDB) reset() {
	db.database = nil
	db.dirty = true
}

func (db *Mb3MongoDB) init() error {
	opt := options.IndexOptions{}
	_, err := db.database.Collection(mbCollection).Indexes().CreateOne(context.Background(),
		mongo.IndexModel{Keys: bson.D{{"accession", 1}},
			Options: opt.SetName("accession_1").SetUnique(true)})
	indices := []string{
		"contributor",
		"compound.names",
		"compound.mass",
		"compound.formula",
		"acquisition.instrumenttype",
		"acquisition.massspectrometry.key",
		"acquisition.massspectrometry.value",
		"peak.peak.intensity",
		"peak.peak.mz",
	}
	for _, index := range indices {
		if _, err := db.database.Collection(mbCollection).Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{Keys: bson.D{{
				index,
				1,
			}},
				Options: &options.IndexOptions{}},
		); err != nil {
			return err
		}
	}
	opt = options.IndexOptions{}
	if _, err := db.database.Collection(mbMetaCollection).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{
				"version",
				1,
			}, {
				"commit",
				1,
			}, {
				"timestamp",
				1,
			}},
			Options: opt.SetUnique(true),
		},
	); err != nil {
		return err
	}
	return err
}

func unmarshal2SearchResult(value bson.M) (*SearchResultData, error) {
	var names = []string{}
	names = getNestedValues(value["names"].(bson.A))
	spectra := []SpectrumMetaData{}
	for _, sp := range value["spectra"].(bson.A) {
		spm := sp.(bson.M)
		s := SpectrumMetaData{spm["id"].(string), spm["title"].(string)}
		spectra = append(spectra, s)
	}

	var result = SearchResultData{
		Names:   names,
		Formula: value["formula"].(bson.A)[0].(string),
		Mass:    value["mass"].(bson.A)[0].(float64),
		Smiles:  value["smiles"].(bson.A)[0].(string),
		Spectra: spectra,
	}
	return &result, nil
}

func getNestedValues(namesAA bson.A) []string {
	var names = []string{}
	for _, namesA := range namesAA {
		for _, name := range namesA.(bson.A) {
			names = append(names, name.(string))
		}
	}
	return removeDuplicateStr(names)
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func unmarshal2Massbank(err error, value *bson.M) (*massbank.MassBank2, error) {
	var mb massbank.MassBank2
	b, err := bson.Marshal(value)
	if err != nil {
		return nil, err
	}
	if err = bson.Unmarshal(b, &mb); err != nil {
		return nil, err
	}

	if mb.Deprecated != nil && reflect.DeepEqual(*mb.Deprecated, massbank.RecordDeprecated{}) {
		mb.Deprecated = nil
	}
	if mb.Project != nil && *mb.Project == "" {
		mb.Project = nil
	}
	if mb.Peak.Annotation != nil && reflect.DeepEqual(*mb.Peak.Annotation, massbank.PkAnnotation{}) {
		mb.Peak.Annotation = nil
	}
	if mb.Species.Name != nil && *mb.Species.Name == "" {
		mb.Species.Name = nil
	}
	if mb.Publication != nil && *mb.Publication == "" {
		mb.Publication = nil
	}
	if mb.Copyright != nil && *mb.Copyright == "" {
		mb.Copyright = nil
	}
	return &mb, err
}
