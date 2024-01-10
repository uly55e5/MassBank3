import './Documentation.scss';

function Documentation() {
  return (
    <div className="documentation-panel">
      <p>
        A documentation is accessible at our{' '}
        {
          <a
            href="https://massbank.github.io/MassBank-documentation/"
            target="_blank"
          >
            GitHub
          </a>
        }{' '}
        repository.
      </p>
    </div>
  );
}

export default Documentation;
