import {createClient} from 'redis'
import {redisConfig} from '../config/redisConfig.js'

export class RedisClientFactory {
    static createClient() {
        return createClient({
            socket: {
                host: redisConfig.host,
                port: redisConfig.port,
            },
        })
    }
}