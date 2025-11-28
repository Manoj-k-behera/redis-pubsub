import { Subscriber } from "../src";
import { expect, jest } from "@jest/globals";

test("Subscriber invoke callback on message", async () => {
    const mockMessage = { msg: "Hello, Subscriber!" };
    const channel = "test-channel";
    const mockclient = {
        connect: jest.fn(),
        subscribe: jest.fn((_, cb) => cb(JSON.stringify(mockMessage))),
        quit: jest.fn(),
    };

    const sub = new Subscriber(mockclient);
    const callback = jest.fn();

    await sub.connect();
    await sub.subscribe(channel, callback);
    await sub.disconnect();

    expect(mockclient.connect).toHaveBeenCalled();
    expect(mockclient.subscribe).toHaveBeenCalledWith(
        channel,
        expect.any(Function)
    );
    expect(callback).toHaveBeenCalledWith(mockMessage);
    expect(mockclient.quit).toHaveBeenCalled();
})