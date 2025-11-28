export class Publisher {
    constructor(redisClient) {
        this.client = redisClient;
    }

    async connect() {
        await this.client.connect();
    }

    async publish(channel, message) {
        await this.client.publish(channel, JSON.stringify(message));
    }

    async disconnect() {
        await this.client.quit();
    }
}