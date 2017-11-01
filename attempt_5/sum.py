import gopy

sum = gopy.load('sum')

print sum.Sum(40, 2)
print sum.Sum(2, 2)

objects = [
    sum.NewSomeObject('John {0}'.format(i)) for i in range(0, 16)
]

for object in objects:
    print object, object.Name
