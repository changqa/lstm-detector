import numpy as np

test = np.load('/home/ensdaddy/packetsdata/data/test_100.npy')
normal = np.load('/home/ensdaddy/packetsdata/data/patch.npy')

print(test[3])
print(normal[3])

def replace(testArray,normalArray,replaceRange):
    for i in replaceRange:
        testArray[i] = normalArray[i]
    return testArray

result = replace(test,normal,[0,1,2,3,4,5,6,7])

print(result[3])

# np.save('/home/ensdaddy/packetsdata/data/modified_test',result)