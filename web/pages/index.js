import fetch from 'isomorphic-fetch';
import Layout from '../components/layout';

const Page = ({ host, data }) => (
  <Layout title="home" host={host}>
    <p>
      Welcome!
    </p>
    <p>
      from api:
    </p>
    <p>
      {JSON.stringify(data)}
    </p>
  </Layout>
);

Page.getInitialProps = async ({ req }) => {
  const apiHost = process.env.API_HOST;
  const endpoint = req ? `http://${apiHost}/graphql` : '/api/graphql';
  const response = await fetch(endpoint, {
    method: 'POST',
    body: JSON.stringify({
      query: 'query {server, secrets { name, value }, token}',
      variables: null
    }),
    headers: { 'Content-Type': 'application/json' }
  });
  const data = await response.json();
  return {
    host: process.env['HOSTNAME'] || (req ? 'localhost' : 'client'),
    data
  };
};

export default Page;
