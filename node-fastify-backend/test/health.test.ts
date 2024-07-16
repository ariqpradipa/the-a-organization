// src/__tests__/app.test.ts
import request from 'supertest';
import fastify from 'fastify';
import app from '../src/app';

describe('GET /health/check', () => {
  it('should return status 200 and a status OK message', async () => {
    const server = fastify();
    server.register(app);
    await server.ready();

    const res = await request(server.server).get('/health/check');
    expect(res.status).toBe(200);
    expect(res.body).toEqual({ status: true });

    await server.close();
  });
});
