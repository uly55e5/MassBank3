import {
  faBarcode,
  faChartColumn,
  faFlask,
  faShareNodes,
  faSignature,
  faSliders,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Form, Input, InputNumber } from 'antd';
import SearchFields from '../../../../types/filterOptions/SearchFields';
import PeakSearch from './searchPanel/peakSearch/PeakSearch';
import PropertyFilterOptionsMenuItems from './searchPanel/msSpecFilter/PropertyFilterOptionsMenuItems';
import StructuralEditor from '../../../basic/StructuralEditor';
import ContentFilterOptions from '../../../../types/filterOptions/ContentFilterOtions';

const peakListPattern =
  /^(\d+(\.\d+)* \d+(\.\d+) \d+(\.\d+)*)(\n\d+(\.\d+)* \d+(\.\d+) \d+(\.\d+)*)*$/;

type InputProps = {
  propertyFilterOptions: ContentFilterOptions | undefined;
  initialStructure: string;
  width: number;
};

function SearchPanelMenuItems({
  propertyFilterOptions,
  initialStructure,
  width,
}: InputProps) {
  return [
    {
      key: 'compoundSearchMenuItem',
      label: 'Compound Search',
      icon: <FontAwesomeIcon icon={faSliders} />,
      children: [
        {
          key: 'compoundName',
          style: {
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            marginLeft: 0,
          },
          label: (
            <Form.Item<SearchFields>
              label="Name"
              name={['compoundSearchFilterOptions', 'compoundName']}
              rules={[{ required: false }]}
              style={{
                width: '100%',
                height: '100%',
              }}
              labelAlign="left"
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <Input type="text" placeholder="Rutin" allowClear />
            </Form.Item>
          ),
        },
        {
          key: 'compoundClass',
          style: {
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            marginLeft: 0,
          },
          label: (
            <Form.Item<SearchFields>
              label="Class"
              name={['compoundSearchFilterOptions', 'compoundClass']}
              rules={[{ required: false }]}
              style={{
                width: '100%',
                height: '100%',
              }}
              labelAlign="left"
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <Input type="text" placeholder="Natural Product" allowClear />
            </Form.Item>
          ),
        },
        {
          key: 'formula',
          style: {
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            marginLeft: 0,
          },
          label: (
            <Form.Item<SearchFields>
              label="Formula"
              name={['compoundSearchFilterOptions', 'formula']}
              rules={[{ required: false }]}
              style={{
                width: '100%',
                height: '100%',
              }}
              labelAlign="left"
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <Input type="text" placeholder="C27H30O16" allowClear />
            </Form.Item>
          ),
        },
        {
          key: 'exactMass',
          style: {
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            marginLeft: 0,
          },
          label: (
            <Form.Item<SearchFields>
              label="Exact Mass"
              name={['compoundSearchFilterOptions', 'exactMass']}
              rules={[{ required: false }]}
              style={{
                width: '100%',
                height: '100%',
              }}
              labelAlign="left"
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <InputNumber placeholder="610.15338" step={0.01} min={0} />
            </Form.Item>
          ),
        },
        {
          key: 'massTolerance_basic',
          style: {
            width: '100%',
            height: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            marginLeft: 0,
          },
          label: (
            <Form.Item<SearchFields>
              label="Mass Tolerance"
              name={['compoundSearchFilterOptions', 'massTolerance']}
              rules={[{ required: false }]}
              style={{
                width: '100%',
                height: '100%',
              }}
              labelAlign="left"
              labelCol={{ span: 8 }}
              wrapperCol={{ span: 16 }}
            >
              <InputNumber placeholder="0.1" step={0.01} min={0} />
            </Form.Item>
          ),
        },
        {
          key: 'inchiMenuItem',
          label: 'InChI',
          icon: <FontAwesomeIcon icon={faSignature} />,
          children: [
            {
              key: 'inchi',
              style: {
                width: '100%',
                height: '100%',
                marginLeft: 0,
              },
              label: (
                <Form.Item<SearchFields>
                  name={['compoundSearchFilterOptions', 'inchi']}
                  rules={[{ required: false }]}
                  style={{
                    width: '100%',
                    height: '100%',
                  }}
                >
                  <Input
                    type="text"
                    placeholder="IKGXIBQEEMLURG-NVPNHPEKSA-N"
                    allowClear
                  />
                </Form.Item>
              ),
            },
          ],
        },
        {
          key: 'structureMenuItem',
          label: 'Structure',
          icon: <FontAwesomeIcon icon={faShareNodes} />,
          children: [
            {
              key: 'structure',
              style: {
                width: '100%',
                height: 700,
                marginLeft: 0,
                overflow: 'scroll',
              },
              label: (
                <StructuralEditor
                  width={width - 100}
                  initialSMILES={initialStructure}
                />
              ),
            },
          ],
        },
      ],
    },
    {
      key: 'spectralSearchMenuItem',
      label: 'Spectral Search',
      icon: <FontAwesomeIcon icon={faChartColumn} />,
      children: [
        {
          key: 'similaritySearchMenuItem',
          label: 'Similarity Search',
          children: [
            {
              key: 'peakList',
              style: {
                width: '100%',
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                marginLeft: 0,
              },
              label: (
                <Form.Item<SearchFields>
                  label="Peak List"
                  name={[
                    'spectralSearchFilterOptions',
                    'similarity',
                    'peakList',
                  ]}
                  rules={[{ required: false, pattern: peakListPattern }]}
                  style={{
                    width: '100%',
                    height: '100%',
                  }}
                  labelAlign="left"
                  labelCol={{ span: 7 }}
                  wrapperCol={{ span: 17 }}
                >
                  <Input.TextArea
                    placeholder="m/z and intensity, delimited by a space. &#10;&#10;147.063 11&#10;303.05 999&#10;449.108 64&#10;465.102 588&#10;611.161 670"
                    autoSize={{ minRows: 5 }}
                    allowClear
                  />
                </Form.Item>
              ),
            },
            {
              key: 'threshold',
              style: {
                width: '100%',
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                marginLeft: 0,
              },
              label: (
                <Form.Item<SearchFields>
                  label="Score Threshold"
                  name={[
                    'spectralSearchFilterOptions',
                    'similarity',
                    'threshold',
                  ]}
                  rules={[{ required: false }]}
                  style={{
                    width: '100%',
                    height: '100%',
                  }}
                  labelAlign="left"
                  labelCol={{ span: 7 }}
                  wrapperCol={{ span: 17 }}
                >
                  <InputNumber placeholder="0.8" step={0.05} min={0} max={1} />
                </Form.Item>
              ),
            },
          ],
        },
        {
          key: 'peakSearchMenuItem',
          label: 'Peak Search',
          children: [
            {
              key: 'peaks_panel',
              style: {
                width: '100%',
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                marginLeft: 0,
              },
              label: <PeakSearch />,
            },
            {
              key: 'massTolerance_peaks',
              style: {
                width: '100%',
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                marginLeft: 0,
              },
              label: (
                <Form.Item<SearchFields>
                  label="Mass Tolerance"
                  name={[
                    'spectralSearchFilterOptions',
                    'peaks',
                    'massTolerance',
                  ]}
                  rules={[{ required: false }]}
                  style={{
                    width: '100%',
                    height: '100%',
                  }}
                  labelAlign="left"
                  labelCol={{ span: 7 }}
                  wrapperCol={{ span: 17 }}
                >
                  <InputNumber placeholder="0.1" step={0.01} min={0} max={1} />
                </Form.Item>
              ),
            },
            {
              key: 'intensity',
              style: {
                width: '100%',
                height: '100%',
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                marginLeft: 0,
              },
              label: (
                <Form.Item<SearchFields>
                  label="Min. Intensity"
                  name={['spectralSearchFilterOptions', 'peaks', 'intensity']}
                  rules={[{ required: false }]}
                  style={{
                    width: '100%',
                    height: '100%',
                  }}
                  labelAlign="left"
                  labelCol={{ span: 7 }}
                  wrapperCol={{ span: 17 }}
                >
                  <InputNumber placeholder="50" step={5} min={0} />
                </Form.Item>
              ),
            },
          ],
        },
        {
          key: 'peakDifferenceMenuItem',
          label: 'Neutral Loss Search',
          children: [
            {
              key: 'neutralLossSearch',
              style: {
                width: '100%',
                height: '100%',
                marginLeft: 0,
                // display: 'flex',
                // justifyContent: 'center',
                // alignItems: 'center',
                textAlign: 'center',
                color: 'orange',
              },
              label: 'Feature is not yet implemented!',
            },
          ],
        },
        {
          key: 'splashMenuItem',
          label: 'SPLASH',
          icon: <FontAwesomeIcon icon={faBarcode} />,
          children: [
            {
              key: 'splash',
              style: {
                width: '100%',
                height: '100%',
                marginLeft: 0,
              },
              label: (
                <Form.Item<SearchFields>
                  name={['spectralSearchFilterOptions', 'splash']}
                  rules={[{ required: false }]}
                  style={{
                    width: '100%',
                    height: '100%',
                  }}
                >
                  <Input
                    type="text"
                    placeholder="splash10-0wmi-0009506000-98ca7f7c8f3072af4481"
                    allowClear
                  />
                </Form.Item>
              ),
            },
          ],
        },
      ],
    },
    {
      key: 'massSpecMenuItem',
      label: 'Property Filter',
      icon: <FontAwesomeIcon icon={faFlask} />,
      children: PropertyFilterOptionsMenuItems({ propertyFilterOptions }),
    },
  ];
}

export default SearchPanelMenuItems;
