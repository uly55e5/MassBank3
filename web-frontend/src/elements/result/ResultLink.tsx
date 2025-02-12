import Hit from '../../types/Hit';
import { Content } from 'antd/es/layout/layout';
import { CSSProperties } from 'react';
import { usePropertiesContext } from '../../context/properties/properties';
import routes from '../../constants/routes';

type InputProps = {
  hit: Hit;
  width?: CSSProperties['width'];
  height?: CSSProperties['height'];
};

function ResultLink({ hit, width = '100%', height = '100%' }: InputProps) {
  const { baseUrl, frontendUrl } = usePropertiesContext();

  const url =
    frontendUrl + baseUrl + routes.accession.path + '?id=' + hit.accession;

  return (
    hit.record && (
      <Content
        style={{
          width,
          height,
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <a
          className="link"
          href={hit.accession && hit.accession !== '' ? url : '?'}
          target="_blank"
        >
          <label style={{ cursor: 'pointer' }}>{hit.accession}</label>
        </a>
      </Content>
    )
  );
}

export default ResultLink;
