import grpc
from concurrent import futures
import time

import game_pb2
import game_pb2_grpc


# 实现 proto 中定义的服务
class GameService(game_pb2_grpc.GameServiceServicer):
    def Login(self, request, context):
        print(f"[Login] 用户名: {request.username}, 密码: {request.password}")
        return game_pb2.LoginResponse(token="test-token-abc123")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    game_pb2_grpc.add_GameServiceServicer_to_server(GameService(), server)
    server.add_insecure_port("[::]:50051")  # 监听本地 50051 端口
    server.start()
    print("gRPC 服务器已启动，监听端口 50051...")
    try:
        while True:
            time.sleep(3600)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    serve()
