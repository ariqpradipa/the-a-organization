import tap from 'tap';
import fastify from 'fastify';
import app from '../../src/app';

tap.test('GET `/health/check` route', async t => {
  // Creating fastify instance
  const server = fastify();

  // Register the app (server with routes) to the fastify instance
  server.register(app);

  // Wait until the server is ready
  await server.ready();

  // Perform the request
  const response = await server.inject({
    method: 'GET',
    url: '/health/check'
  });

  // Assertions
  t.equal(response.statusCode, 200, 'Response status should be 200');
  t.same(response.json(), { status: true }, 'Response should be { status: "true" }');

  t.end();
});
