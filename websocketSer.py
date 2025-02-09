import asyncio

import websocket

import logging 



logging.basicConfig(level=logging.INFO)


async def heartbeat(websocket):
    " ping & pong"
    while True:
        try:
            logging.info("Sending ping")
            # 发送ping，并等待pong（pong_waiter 为Future对象）
            pong_waiter = await websocket.ping()
            await 