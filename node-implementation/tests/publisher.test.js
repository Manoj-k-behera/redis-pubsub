import { Publisher } from "../src";
import { jest } from "@jest/globals";

test("Publisher publishes messages without any errors", async () => {
    const mockclient = {
        connect: jest.fn(),
        publish: jest.fn(),
        quit: jest.fn(),
    };

    const channel = "test-channel";
    const mockMessage = { msg: "Hello, World!" };

    const pub = new Publisher(mockclient);

    await pub.connect();
    await pub.publish(channel, mockMessage);
    await pub.disconnect();

    expect(mockclient.connect).toHaveBeenCalled();
    expect(mockclient.publish).toHaveBeenCalledWith(
        channel,
        JSON.stringify(mockMessage)
    );
    expect(mockclient.quit).toHaveBeenCalled();
});
