final = [1]

def TriangleNumber(l):
    for i in range(1,l+1):
        final.append(final[-1]+1+i)


TriangleNumber(6)
print final
