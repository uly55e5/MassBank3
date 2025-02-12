import { Button, Input } from 'antd';
import { Content } from 'antd/es/layout/layout';
import {
  ChangeEvent,
  CSSProperties,
  KeyboardEvent,
  useCallback,
  useState,
} from 'react';
import { createSearchParams, useNavigate } from 'react-router-dom';
import routes from '../../constants/routes';
import { usePropertiesContext } from '../../context/properties/properties';

type InputProps = {
  width: CSSProperties['width'];
  height: CSSProperties['height'];
  accession?: string;
};

function AccessionSearchInputField({
  width,
  height,
  accession: acc,
}: InputProps) {
  const [accession, setAccession] = useState<string>(acc ?? '');
  const navigate = useNavigate();
  const { baseUrl } = usePropertiesContext();

  const handleOnChange = useCallback((e: ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    e.stopPropagation();

    setAccession(e.target.value.trim());
  }, []);

  const handleOnClick = useCallback(
    () =>
      navigate({
        pathname: baseUrl + routes.accession.path,
        search: `?${createSearchParams({ id: accession })}`,
      }),
    [accession, baseUrl, navigate],
  );

  const handleOnKeyDown = useCallback(
    (e: KeyboardEvent<HTMLInputElement>) => {
      if (e.key === 'Enter') {
        handleOnClick();
      }
    },
    [handleOnClick],
  );

  return (
    <Content
      style={{
        width,
        height,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#f3ece0',
      }}
    >
      <Input
        type="text"
        placeholder="e.g. MSBNK-AAFC-AC000114"
        value={accession && accession !== '' ? accession : undefined}
        addonBefore="Go to accession:"
        onChange={handleOnChange}
        onKeyDown={handleOnKeyDown}
        allowClear
        style={{ width: 500 }}
      />
      <Button
        children="Search"
        onClick={handleOnClick}
        disabled={accession.trim() === ''}
        style={{ width: 100, marginLeft: 20 }}
      />
    </Content>
  );
}

export default AccessionSearchInputField;
