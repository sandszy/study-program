import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
plt.rcParams["font.sans-serif"]=["SimHei"] #设置字体
plt.rcParams["axes.unicode_minus"]=False #该语句解决图像中的“-”负号的乱码问题

region = pd.read_excel(io="./cloud_resource_pool_testdata.xlsx")

region["Saturability"] = region[["CPU allocation rate","MEM allocation rate"]].max(axis=1)

region["allocated"] = region["Number of nodes"] * region["Saturability"]

region["unallocated"] = region["Number of nodes"] - region["allocated"]

AllocationTable = region[["Region name","allocated","unallocated"]]

# 将 "Region name" 列设置为索引
AllocationTable.set_index("Region name", inplace=True)

# 画堆叠图
AllocationTable.plot(kind='bar', stacked=True)

# 添加图例和标签
plt.legend(title='Allocation')
plt.xlabel('Region')
plt.ylabel('Values')
plt.title('Stacked Bar Chart - Allocation Table')

# 显示图形
plt.show()
# # Plotting
# fig, ax1 = plt.subplots(figsize=(10, 6))

# # Bar plot for CPU allocation rate
# ax1.bar(CPU["Region name"], CPU["CPU allocation rate"], color='b', alpha=0.7, label='CPU allocation rate')
# ax1.set_xlabel('Region name', fontsize=14)
# ax1.set_ylabel('CPU allocation rate', fontsize=14, color='b')
# ax1.tick_params('y', colors='b')
# ax1.legend(loc='upper left')

# # Create a second y-axis for Number of CPUs
# ax2 = ax1.twinx()
# ax2.plot(CPU["Region name"], CPU["Number of CPUs"], color='r', marker='o', label='Number of CPUs')
# ax2.set_ylabel('Number of CPUs', fontsize=14, color='r')
# ax2.tick_params('y', colors='r')
# ax2.legend(loc='upper right')

# # Title and show the plot
# plt.title('CPU Allocation and Number of CPUs by Region', fontsize=16)
# plt.show()