import { RedisClientFactory, Publisher, Subscriber } from "../src/index.js";

test("End-to-end Pub/Sub test", async () => {
  const pubClient = RedisClientFactory.createClient();
  const subClient = RedisClientFactory.createClient();

  const publisher = new Publisher(pubClient);
  const subscriber = new Subscriber(subClient);

  await publisher.connect();
  await subscriber.connect();
  
  const channel = "demo-channel";
  const testMessage = { data: 123 };

  const messageReceived = new Promise((resolve) => {
    subscriber.subscribe(channel, (msg) => resolve(msg));
  });

  await publisher.publish(channel, testMessage);

  const result = await messageReceived;

  expect(result).toEqual(testMessage);

  await publisher.disconnect();
  await subscriber.disconnect();
});
