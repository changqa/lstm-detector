
def getAttackIndexArry(filename):
    result = []
    with open(filename, 'r+', encoding='utf-8') as f:
        for line in f.readlines():
            oneattack = []
            currentAttack = line.strip().split(" ")
            oneattack.append(int(currentAttack[0]))
            oneattack.append(int(currentAttack[-1]))
            result.append(oneattack)
        return result

index = getAttackIndexArry("/home/ensdaddy/packetsdata/test/attack_index")
result = []
chosen = [1,8,10,11,16,18,20,21,24,26,28,32,34,36,40,41,48,50,51,64]
for i in chosen:
    result.append(index[i])
print(result)
con = []
for i in range(len(result)):
    con.append('contextual')
print(con)