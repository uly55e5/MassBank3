import './Table.scss';

import { Content } from 'antd/es/layout/layout';
import ExportableContent from '../common/ExportableContent';
import { CSSProperties, JSX, useCallback, useMemo } from 'react';
import copyTextToClipboard from '../../utils/copyTextToClipboard';
import routes from '../../constants/routes';
import { Table } from 'antd';
import Record from '../../types/Record';
import { MF } from 'react-mf';
import StructureView from '../basic/StructureView';
import LabelWrapper from './LabelWrapper';
import { usePropertiesContext } from '../../context/properties/propertiesContext';

const titleHeight = 50;
const labelWidth = 120;

type HeaderTableType = {
  key: string;
  label: string;
  value: JSX.Element;
  copyButton?: JSX.Element;
};
const columns = [
  {
    dataIndex: 'label',
    key: 'record-view-header-table-label',
    width: labelWidth,
    align: 'left' as const,
  },
  {
    dataIndex: 'value',
    key: 'record-view-header-table-value',
    width: `calc(100% - ${labelWidth})`,
    align: 'left' as const,
  },
];

type InputProps = {
  record: Record;
  width: CSSProperties['width'];
  height: CSSProperties['height'];
  imageWidth: CSSProperties['width'];
};

function RecordViewHeader({ record, width, height, imageWidth }: InputProps) {
  const { baseUrl, frontendUrl } = usePropertiesContext();

  const handleOnCopy = useCallback((label: string, text: string) => {
    copyTextToClipboard(label, text);
  }, []);

  const buildSearchUrl = useCallback(
    (label: string, value: string) => {
      const searchParams = new URLSearchParams();
      searchParams.set(label, value);
      const url =
        frontendUrl +
        baseUrl +
        routes.search.path +
        `?${searchParams.toString()}`;

      return url;
    },
    [baseUrl, frontendUrl],
  );

  return useMemo(() => {
    const dataSource: HeaderTableType[] = [];
    dataSource.push({
      key: 'record-view-header-table-names',
      label: 'Names',
      value: (
        <Content
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'center',
            alignItems: 'left',
          }}
        >
          {record.compound.names.map((name, i) => (
            <ExportableContent
              key={'name-label-' + name}
              component={<LabelWrapper value={name} />}
              mode="copy"
              onClick={() => handleOnCopy(`Compound name ${i + 1}`, name)}
              title={`Copy compound name ${i + 1} to clipboard`}
              enableSearch
              searchTitle={`Search for compound name ${i + 1}`}
              searchUrl={buildSearchUrl('compound_name', name)}
            />
          ))}
        </Content>
      ),
    });
    dataSource.push({
      key: 'record-view-header-table-classes',
      label: 'Classes',
      value:
        record.compound.classes &&
        record.compound.classes.length === 1 &&
        record.compound.classes[0] !== 'N/A' ? (
          <ExportableContent
            component={<LabelWrapper value={record.compound.classes[0]} />}
            mode="copy"
            onClick={() =>
              handleOnCopy('Compound classes', record.compound.classes[0])
            }
            title="Copy compound classes to clipboard"
          />
        ) : (
          <label style={{ color: 'grey', fontStyle: 'italic' }}>N/A</label>
        ),
    });
    dataSource.push({
      key: 'record-view-header-table-smiles',
      label: 'SMILES',
      value: (
        <ExportableContent
          component={<LabelWrapper value={record.compound.smiles} />}
          mode="copy"
          onClick={() => handleOnCopy('SMILES', record.compound.smiles)}
          title="Copy SMILES to clipboard"
          enableSearch
          searchTitle="Search for SMILES"
          searchUrl={buildSearchUrl('substructure', record.compound.smiles)}
        />
      ),
    });
    dataSource.push({
      key: 'record-view-header-table-inchi',
      label: 'InChI',
      value: (
        <ExportableContent
          component={<LabelWrapper value={record.compound.inchi} />}
          mode="copy"
          onClick={() => handleOnCopy('InChI', record.compound.inchi)}
          title="Copy InChi to clipboard"
          enableSearch
          searchTitle="Search for InChI"
          searchUrl={buildSearchUrl('inchi', record.compound.inchi)}
        />
      ),
    });
    dataSource.push({
      key: 'record-view-header-table-splash',
      label: 'SPLASH',
      value: (
        <ExportableContent
          component={record.peak.splash}
          mode="copy"
          onClick={() => handleOnCopy('SPLASH', record.peak.splash)}
          title="Copy SPLASH to clipboard"
          enableSearch
          searchTitle="Search for SPLASH"
          searchUrl={buildSearchUrl('splash', record.peak.splash)}
        />
      ),
    });

    return (
      <Content
        style={{
          width,
          minHeight: height,
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          backgroundColor: 'white',
        }}
      >
        <ExportableContent
          component={<LabelWrapper value={record.title} />}
          componentContainerStyle={{
            minHeight: titleHeight,
            maxHeight: titleHeight,
            fontSize: 18,
            fontWeight: 'bold',
            width: '100%',
            justifyContent: 'center',
          }}
          mode="copy"
          onClick={() => handleOnCopy('Title', record.title)}
          title="Copy title to clipboard"
        />
        <Content
          style={{
            width: '100%',
            minHeight: `calc(${height} - ${titleHeight})`,
            maxHeight: `calc(${height} - ${titleHeight})`,
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <Table<HeaderTableType>
            style={{
              minWidth: `calc(100% - ${imageWidth})`,
              maxWidth: `calc(100% - ${imageWidth})`,
              height: '100%',
            }}
            className="table"
            sticky
            pagination={false}
            showHeader={false}
            columns={columns}
            dataSource={dataSource}
          />
          <Content
            style={{
              minWidth: imageWidth,
              maxWidth: imageWidth,
              height: '100%',
              display: 'flex',
              flexDirection: 'column',
              justifyContent: 'center',
              alignItems: 'center',
            }}
          >
            {record.compound.smiles && record.compound.smiles !== '' ? (
              <StructureView
                smiles={record.compound.smiles}
                imageWidth={imageWidth as number}
                imageHeight={(height as number) - titleHeight - 80}
              />
            ) : undefined}

            <Content
              style={{
                width: imageWidth,
                height: `calc(100% - ${height} - ${titleHeight} - 80px)`,
                display: 'grid',
                gridTemplateColumns: '75px auto 55px auto',
                alignItems: 'center',
                paddingTop: 10,
              }}
            >
              <label>Formula: </label>
              <ExportableContent
                component={<MF mf={record.compound.formula} />}
                componentContainerStyle={{
                  fontWeight: 'bolder',
                }}
                mode="copy"
                onClick={() =>
                  copyTextToClipboard(
                    'Molecular Formula',
                    record.compound.formula,
                  )
                }
                title="Copy molecular formula to clipboard"
                enableSearch
                searchTitle="Search for molecular formula"
                searchUrl={buildSearchUrl('formula', record.compound.formula)}
              />
              <label>Mass: </label>
              <ExportableContent
                component={record.compound.mass.toString()}
                componentContainerStyle={{
                  fontWeight: 'bolder',
                }}
                mode="copy"
                onClick={() =>
                  copyTextToClipboard(
                    'Molecular Mass',
                    record.compound.mass.toString(),
                  )
                }
                title="Copy molecular mass to clipboard"
                enableSearch
                searchTitle="Search for molecular mass"
                searchUrl={buildSearchUrl(
                  'exact_mass',
                  record.compound.mass.toString(),
                )}
              />
            </Content>
          </Content>
        </Content>
      </Content>
    );
  }, [
    buildSearchUrl,
    handleOnCopy,
    height,
    imageWidth,
    record.compound.classes,
    record.compound.formula,
    record.compound.inchi,
    record.compound.mass,
    record.compound.names,
    record.compound.smiles,
    record.peak.splash,
    record.title,
    width,
  ]);
}

export default RecordViewHeader;
