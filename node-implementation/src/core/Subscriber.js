export class Subscriber {
    constructor(redisClient) {
        this.client = redisClient;
    }

    async connect() {
        await this.client.connect();
    }

    async subscribe(channel, messageHandler) {
        await this.client.subscribe(channel, (message) => {
            const parsedMessage = JSON.parse(message);
            messageHandler(parsedMessage);
        });
    }

    async disconnect() {
        await this.client.quit();
    }
}
