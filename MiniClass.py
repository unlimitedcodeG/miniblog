class LoggingMixin:
    def log(self, message):
        print(f"Log:{message}")

    def log_error(self, message):
        print(f"Error:{message}")


class Db:
    def __init__(self, data):
        self.data = data

    def add(self, item):
        self.data.append(item)
        self.log("Item added")
