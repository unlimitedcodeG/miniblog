import asyncio
import websockets


async def chat():
    uri = "ws://localhost:6790"
    async with websockets.connect(uri) as websocket:
        print("输入 'exit' 退出聊天")
        while True:
            message = input("请输入消息: ")
            if message == "exit":
                break
            # 发送消息到服务器
            await websocket.send(message)
            # 接收来自服务器的消息
            response = await websocket.recv()
            print(f"收到: {response}")


asyncio.run(chat())
