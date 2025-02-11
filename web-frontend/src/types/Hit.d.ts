export default interface Hit {
  index: number;
  accession: string;
  atomcount: number;
  score?: number;
  record?: Record;
}
