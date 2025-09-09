const redis = require('redis');

// Read host and port from environment variables, with defaults
const redisHost = process.env.REDIS_HOST || 'localhost';
const redisPort = process.env.REDIS_PORT || '6379';

// Construct the Redis URL
const redisUrl = `redis://${redisHost}:${redisPort}`;

const client = redis.createClient({ url: redisUrl });

client.on('error', (err) => console.error('Redis Client Error:', err));

client.connect()
  .then(() => console.log(`✅ Connected to Redis at ${redisUrl}`))
  .catch((err) => console.error('❌ Redis connection failed:', err));

module.exports = client;