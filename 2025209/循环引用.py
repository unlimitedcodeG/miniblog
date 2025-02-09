import weakref


class Node:
    def __init__(self, value):
        self.value = value
        self.parent = None
        self.children = []

    def add_child(self, child):
        child.parent = weakref(self)
        self.children.append(child)


parent = Node("parent")
child = Node("child")
parent.add_child(child)

# Python 中避免循环引用问题的几个方法可以帮助你更有效地管理内存。
# 循环引用通常发生在两个或更多对象互相持有对方的引用，
# 而这可能导致 Python 的垃圾回收机制无法回收这些对象。
# 以下是一些常用的方法：
# 使用弱引用： Python 的 weakref 模块可以用来创建对象的弱引用。
# # 弱引用不会增加对象的引用计数，因此不会阻止垃圾回收器回收引用的对象。
# # 如果你的对象之间存在潜在的循环引用，可以考虑将其中一个引用改为弱引用。


# Python 中避免循环引用问题的几个方法可以帮助你更有效地管理内存。循环引用通常发生在两个或更多对象互相持有对方的引用，
# 而这可能导致 Python 的垃圾回收机制无法回收这些对象。以下是一些常用的方法：

# 使用弱引用： Python 的 weakref
# 模块可以用来创建对象的弱引用。弱引用不会增加对象的引用计数，
# 因此不会阻止垃圾回收器回收引用的对象。如果你的对象之间存在潜在的循环引用，可以考虑将其中一个引用改为弱引用。
parent.children.clear()
child.parent = None
# 使用 gc 模块： Python 标准库中的 gc 模块可以用来强制进行垃圾回收
# 也可以用来识别循环引用中的对象。通过 gc.collect()
# 可以强制立即进行垃圾回收。此外，你可以使用 gc.get_objects() 查看当前存在的对象，帮助你分析内存问题。
import gc

gc.collect()  # 强制回收垃圾
### 
