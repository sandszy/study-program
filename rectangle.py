import turtle

# 创建一个 turtle 对象
t = turtle.Turtle()

# 设置矩形的边长和宽度
length = 200
width = 100

# 绘制矩形
for _ in range(4):
    t.forward(length)  # 向前移动
    t.left(90)         # 向左旋转90度

# 隐藏turtle并显示绘图窗口
turtle.hideturtle()
turtle.done()