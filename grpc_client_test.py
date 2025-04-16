import grpc
import game_pb2
import game_pb2_grpc

# 创建连接到服务端
channel = grpc.insecure_channel("localhost:50051")
stub = game_pb2_grpc.GameServiceStub(channel)

# 构造请求并调用
request = game_pb2.LoginRequest(username="player1", password="123456")
response = stub.Login(request)

# 输出结果
print("[客户端] 接收到 token:", response.token)
