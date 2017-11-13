from threading import Thread

import gopy

sum = gopy.load('sum')

print sum.DoThing()

# def do_some_things(i):
#     print 'starting', i
#     print sum.Sleep()
#     print 'finished', i
#
# # print sum.Sum(40, 2)
#
# # objects = [
# #     sum.NewSomeObject('John {0}'.format(i)) for i in range(0, 16)
# # ]
#
# # for object in objects:
# #     print object, object.Name
# #     for thing in object.Things:
# #         print thing
# #
# # print sum.Sum(2, 2)
#
# threads = []
# for i in range(0, 32):
#     threads += [Thread(
#         target=do_some_things,
#         args=(i,),
#     )]
#
#     threads[-1].start()
#
# for thread in threads:
#     print 'joining', thread
#     thread.join()
