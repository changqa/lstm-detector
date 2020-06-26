import numpy as np 
from matplotlib import pyplot as plt 
import matplotlib

train_np_array = np.load("/home/ensdaddy/origin-lstm/telemanom/data/train/A-1.npy")
test_np_array = np.load("/home/ensdaddy/origin-lstm/telemanom/data/test/A-1.npy")
smoothed_error_np_array = np.load("/home/ensdaddy/origin-lstm/telemanom/data/2020-06-22_08.38.23/smoothed_errors/A-1.npy")
y_hat_np_array = np.load('/home/ensdaddy/origin-lstm/telemanom/data/2020-06-22_08.38.23/y_hat/A-1.npy')

print(len(train_np_array))


x = [i for i in range(12600,12900)]
y = [i[0] for i in train_np_array[12600:12900]]

plt.plot(x,y)
plt.show()