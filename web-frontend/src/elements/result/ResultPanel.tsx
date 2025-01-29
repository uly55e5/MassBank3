import { MouseEvent, useCallback, useEffect, useMemo, useState } from 'react';
import ResultTable from './ResultTable';
import Hit from '../../types/Hit';
import Peak from '../../types/peak/Peak';
import Record from '../../types/Record';
import generateID from '../../utils/generateID';
import Placeholder from '../basic/Placeholder';
import fetchData from '../../utils/request/fetchData';
import {
  Button,
  Dropdown,
  MenuProps,
  Modal,
  Pagination,
  Select,
  Spin,
} from 'antd';
import { Content } from 'antd/es/layout/layout';
import SpectralHitsCarouselView from '../routes/pages/search/SpectralHitsCarouselView';
import ResultTableSortOptionType from '../../types/ResultTableSortOptionType';
import axios from 'axios';
import FileSaver from 'file-saver';
const { saveAs } = FileSaver;

type InputProps = {
  reference?: Peak[];
  hits: Hit[];
  width: number;
  height: number;
  sortOptions?: ResultTableSortOptionType[];

  onSort?: (value: string) => void;
  widthOverview?: number;
  heightOverview?: number;
};

function ResultPanel({
  reference,
  hits,
  width,
  height,
  sortOptions = [],
  onSort = () => {},
  widthOverview = width,
  heightOverview = height,
}: InputProps) {
  const [isRequesting, setIsRequesting] = useState<boolean>(false);
  const [showModal, setShowModal] = useState<boolean>(false);
  const [slideIndex, setSlideIndex] = useState<number>(0);
  const [resultPageIndex, setResultPageIndex] = useState<number>(0);
  const [selectedSortOption, setSelectedSortOption] = useState<
    string | undefined
  >();
  const [hitsWithRecords, setHitsWithRecords] = useState<Hit[] | undefined>();

  const pageLimit = 20;
  const paginationHeight = 50;

  const resultTableData = useMemo(() => {
    const _resultTableData: Hit[][] = [];
    let counter = 0;
    let resultHits: Hit[] = [];

    for (let i = 0; i < hits.length; i++) {
      if (counter < pageLimit) {
        resultHits.push(hits[i]);
        counter++;
      } else {
        _resultTableData.push(resultHits);
        resultHits = [hits[i]];
        counter = 1;
      }
    }
    if (resultHits.length > 0) {
      _resultTableData.push(resultHits);
    }

    return _resultTableData;
  }, [hits]);

  const fetchRecords = useCallback(async () => {
    setIsRequesting(true);

    const _hits =
      resultTableData.length > 0 ? resultTableData[resultPageIndex] : [];

    let _hitsWithRecords: Hit[] = [];
    if (_hits.length > 0) {
      const from = 0;
      let to = pageLimit;
      if (to > _hits.length) {
        to = _hits.length;
      }

      const range = to - from;
      const accessions = _hits
        .slice(from, from + range)
        .map((h) => h.accession);

      const records: (Record | undefined)[] = [];
      for (const accession of accessions) {
        const url =
          process.env.REACT_APP_MB3_API_URL +
          '/v1/records/' +
          accession +
          '/simple';

        const record = await fetchData(url);

        if (record) {
          record.peak.peak.values = record.peak.peak.values.map((p) => ({
            ...p,
            id: generateID(),
          }));
          records.push(record);
        } else {
          records.push(undefined);
        }
      }

      _hitsWithRecords = _hits.slice(from, from + range).map((h, i) => {
        h.record = records[i];
        return h;
      });
    }

    setHitsWithRecords(_hitsWithRecords);
    setIsRequesting(false);
  }, [resultTableData, resultPageIndex]);

  useEffect(() => {
    fetchRecords();
  }, [fetchRecords]);

  const handleOnDoubleClick = useCallback(
    (_slideIndex: number) => {
      setSlideIndex(_slideIndex % pageLimit);
      setShowModal(true);
    },
    [setShowModal],
  );

  const resultTable = useMemo(
    () => (
      <ResultTable
        reference={reference}
        hits={hitsWithRecords || []}
        height={height - paginationHeight}
        onDoubleClick={handleOnDoubleClick}
        rowHeight={150}
        chartWidth={250}
        imageWidth={250}
      />
    ),
    [reference, hitsWithRecords, height, handleOnDoubleClick],
  );

  const modal = useMemo(
    () => (
      <Modal
        open={showModal}
        onCancel={() => setShowModal(false)}
        footer={null}
        width={widthOverview}
        height={heightOverview}
        centered
        destroyOnClose
      >
        <SpectralHitsCarouselView
          reference={reference}
          hits={hitsWithRecords || []}
          slideIndex={slideIndex}
          width={widthOverview}
          height={heightOverview}
        />
      </Modal>
    ),
    [
      showModal,
      widthOverview,
      heightOverview,
      reference,
      hitsWithRecords,
      slideIndex,
    ],
  );

  const handleOnSelectPage = useCallback(
    (pageIndex: number | null) => {
      if (
        pageIndex &&
        pageIndex > 0 &&
        pageIndex <= Math.ceil(hits.length / pageLimit)
      ) {
        setResultPageIndex(pageIndex - 1);
      }
    },
    [hits.length],
  );

  const handleOnSelect = useCallback(
    (value: string) => {
      setSelectedSortOption(value);
      onSort(value);
    },
    [onSort],
  );

  const handleOnDownloadResult = useCallback(
    async (format: string) => {
      setIsRequesting(true);
      const host = process.env.REACT_APP_EXPORT_SERVICE_URL;
      const url = `${host}/convert`;

      const resp = await axios.post(
        url,
        {
          record_list: hits.map((h) => h.accession),
          format,
        },
        {
          headers: {
            'Content-Type': 'application/json',
            Accept: 'application/octet-stream',
          },
        },
      );
      if (resp.status === 200) {
        const data = await resp.data;
        const fileType = format.split('_')[1];
        const filename = `massbank_result.${format}.${fileType}`;
        const blob = new Blob([data], {
          type: 'application/octet-stream',
        });
        saveAs(blob, filename);
      }

      setIsRequesting(false);
    },
    [hits],
  );

  const buildDownloadOptionLabel = useCallback(
    (label: string, format: string) => (
      <label
        onClick={(e: MouseEvent<HTMLLabelElement>) => {
          e.preventDefault();
          handleOnDownloadResult(format);
        }}
        style={{
          width: 100,
          marginLeft: 10,
          marginRight: 10,
        }}
      >
        {label}
      </label>
    ),
    [handleOnDownloadResult],
  );

  const items: MenuProps['items'] = useMemo(
    () => [
      {
        key: '0_nist_msp_download',
        label: buildDownloadOptionLabel('NIST MSP', 'nist_msp'),
      },
      {
        key: '1_riken_msp_download',
        label: buildDownloadOptionLabel('RIKEN MSP', 'riken_msp'),
      },
    ],
    [buildDownloadOptionLabel],
  );

  const paginationContainer = useMemo(() => {
    return (
      <Content
        style={{
          width: '100%',
          height: paginationHeight,
          display: 'flex',
          justifyContent: 'space-evenly',
          alignItems: 'center',
        }}
      >
        <Pagination
          total={hits.length}
          pageSize={pageLimit}
          showTotal={(total) => (
            <Content
              style={{
                textWrap: 'nowrap',
                textAlign: 'center',
                fontWeight: 'bold',
                color: 'brown',
                width: 200,
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
              }}
            >{`${total} Results`}</Content>
          )}
          onChange={handleOnSelectPage}
          current={resultPageIndex + 1}
          showTitle
          showSizeChanger={false}
          showQuickJumper
          locale={{ jump_to: 'Page', page: '' }}
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'space-around',
            alignItems: 'center',
          }}
        />
        {sortOptions.length > 0 && (
          <Select
            defaultValue={selectedSortOption}
            style={{ width: 200 }}
            placeholder="Sort by"
            optionFilterProp="label"
            filterSort={(optionA, optionB) =>
              (optionA?.label ?? '')
                .toLowerCase()
                .localeCompare((optionB?.label ?? '').toLowerCase())
            }
            options={sortOptions}
            onSelect={handleOnSelect}
          />
        )}

        <Dropdown menu={{ items }} trigger={['click']}>
          <Button
            style={{
              width: 100,
              marginRight: 30,
              marginLeft: 20,
            }}
          >
            Download
          </Button>
        </Dropdown>
      </Content>
    );
  }, [
    hits.length,
    handleOnSelectPage,
    resultPageIndex,
    sortOptions,
    selectedSortOption,
    handleOnSelect,
    items,
  ]);

  return useMemo(
    () =>
      resultTableData.length > 0 ? (
        <Content style={{ width, height }}>
          {isRequesting ? (
            <Content
              style={{
                width: '100%',
                height,
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
              }}
            >
              <Spin size="large" />
            </Content>
          ) : (
            <Content
              style={{
                width: '100%',
                height,
                overflow: 'scroll',
              }}
            >
              {paginationContainer}
              {resultTable}
              {modal}
            </Content>
          )}
        </Content>
      ) : (
        <Placeholder
          child={'No results'}
          style={{ width, height, fontSize: 18, fontWeight: 'bold' }}
        />
      ),
    [
      height,
      isRequesting,
      modal,
      paginationContainer,
      resultTable,
      resultTableData.length,
      width,
    ],
  );
}

export default ResultPanel;
