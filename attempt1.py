import pprint

from src.attempt1 import GetOneWithTheLot

one_with_the_lot = GetOneWithTheLot()

print one_with_the_lot

pprint.pprint({
    k: getattr(one_with_the_lot, k) for k in [
        x for x in dir(one_with_the_lot) if not x.startswith('_')
    ]
})
