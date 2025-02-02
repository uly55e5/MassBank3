import Hit from '../../types/Hit';
import Peak from '../../types/peak/Peak';

import { useCallback, useMemo } from 'react';
import { Table } from 'antd';
import ResultTableDataType from '../../types/ResultTableDataType';
import ResultLink from './ResultLink';
import Chart from '../basic/Chart';
import StructureView from '../basic/StructureView';
import { Content } from 'antd/es/layout/layout';

type InputProps = {
  reference?: Peak[];
  hits: Hit[];
  height: number;
  onDoubleClick: (slideIndex: number) => void;
  rowHeight?: number;
  chartWidth?: number;
  imageWidth?: number;
};

function ResultTable({
  reference,
  hits,
  height,
  onDoubleClick,
  rowHeight = 100,
  chartWidth = 200,
  imageWidth = 200,
}: InputProps) {
  const buildChart = useCallback(
    (hit: Hit) =>
      reference && reference.length > 0 ? (
        <Content
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          {' '}
          <Chart
            peakData={reference}
            peakData2={
              (hit.record ? hit.record.peak.peak.values : []) as Peak[]
            }
            width={chartWidth}
            height={rowHeight}
            disableZoom
            disableLabels
            disableOnHover
          />
        </Content>
      ) : (
        <Content
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <Chart
            peakData={(hit.record ? hit.record.peak.peak.values : []) as Peak[]}
            width={chartWidth}
            height={rowHeight}
            disableZoom
            disableLabels
            disableOnHover
          />
        </Content>
      ),
    [chartWidth, reference, rowHeight],
  );

  const buildStructure = useCallback(
    (smiles: string) => (
      <Content
        style={{
          width: '100%',
          height: '100%',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <StructureView
          smiles={smiles}
          imageWidth={imageWidth}
          imageHeight={rowHeight}
          disableExport
        />
      </Content>
    ),
    [imageWidth, rowHeight],
  );

  const dataSource: ResultTableDataType[] = useMemo(() => {
    const rows: ResultTableDataType[] = [];

    hits.forEach((hit) => {
      const row: ResultTableDataType = {
        key: 'result-table-row_' + hit.index + '_' + hit.score,
        index: hit.index + 1,
        score: hit.score ? hit.score.toFixed(4) : undefined,
        accession: <ResultLink hit={hit} />,
        title: hit.record.title,
        chart: buildChart(hit),
        structure: buildStructure(hit.record.compound.smiles),
      };

      rows.push(row);
    });

    return rows;
  }, [buildChart, buildStructure, hits]);

  const handleOnDoubleClick = useCallback(
    (record: ResultTableDataType) => ({
      onDoubleClick: () => {
        onDoubleClick(record.index - 1);
      },
    }),
    [onDoubleClick],
  );

  const columns = useMemo(() => {
    const defaultColumns = [
      {
        title: 'Index',
        dataIndex: 'index',
        key: 'index',
        align: 'center' as const,
        width: 100,
      },
      {
        title: 'Accession',
        dataIndex: 'accession',
        key: 'accession',
        align: 'center' as const,
      },
      {
        title: 'Title',
        dataIndex: 'title',
        key: 'title',
        align: 'center' as const,
      },
      {
        title: 'Chart',
        dataIndex: 'chart',
        key: 'chart',
      },
      {
        title: 'Structure',
        dataIndex: 'structure',
        key: 'structure',
      },
    ];

    const _columns = [...defaultColumns];

    if (hits.find((hit) => hit.score !== undefined)) {
      _columns.splice(1, 0, {
        title: 'Score',
        dataIndex: 'score',
        key: 'score',
        align: 'center' as const,
        width: 100,
      });
    }

    return _columns;
  }, [hits]);

  return useMemo(
    () => (
      <Table<ResultTableDataType>
        style={{
          width: '100%',
          height,
          overflow: 'scroll',
        }}
        columns={columns}
        dataSource={dataSource}
        pagination={false}
        onRow={handleOnDoubleClick}
        sticky
      />
    ),
    [columns, dataSource, handleOnDoubleClick, height],
  );
}

export default ResultTable;
