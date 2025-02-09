import asyncio
import websockets
import logging

logging.basicConfig(level=logging.INFO)


async def handler(websocket):
    logging.info(f"Connected via {websocket}")
    try:
        async for message in websocket:
            logging.info(f"Received message: {message}")
            await websocket.send(f"Echo: {message}")
            logging.info(f"Sent message: Echo: {message}")
    except Exception as e:
        logging.error(f"Error: {e}")
    finally:
        logging.info("Connection closed")


async def main():
    try:
        # Here, ensure that the handler function accepts two parameters: websocket and path.
        async with websockets.serve(handler, "localhost", 6790):
            logging.info("WebSocket server is running at ws://localhost:6790")
            await asyncio.Future()  # This will keep the server running indefinitely.
    except Exception as e:
        logging.error(f"Server error: {e}")


if __name__ == "__main__":
    asyncio.run(main())
