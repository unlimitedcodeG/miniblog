from locust import User, task, between
import grpc
import time
import game_pb2
import game_pb2_grpc


class GrpcUser(User):
    wait_time = between(0.5, 1.5)  # 模拟用户每隔 0.5-1.5 秒发一次请求

    def on_start(self):
        # 建立 gRPC channel 和 stub
        self.channel = grpc.insecure_channel("127.0.0.1:50051")
        self.stub = game_pb2_grpc.GameServiceStub(self.channel)

    @task
    def login(self):
        start_time = time.time()
        try:
            request = game_pb2.LoginRequest(username="user1", password="123456")
            response = self.stub.Login(request)
            elapsed = int((time.time() - start_time) * 1000)

            # 上报成功
            self.environment.events.request_success.fire(
                request_type="grpc",
                name="Login",
                response_time=elapsed,
                response_length=len(response.token),
            )
        except Exception as e:
            elapsed = int((time.time() - start_time) * 1000)
            # 上报失败
            self.environment.events.request_failure.fire(
                request_type="grpc",
                name="Login",
                response_time=elapsed,
                response_length=0,
                exception=e,
            )
