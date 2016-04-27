import numpy
import theano
import theano.tensor as T

# rng = numpy.random

nfields = None
opcount = []
size = []
with open("costs.txt") as f:
	for line in f.readlines():
		fields = line.split()
		if nfields is None:
			# header line
			nfields = len(fields)
			ops = fields
			print(line)
			continue
		elif nfields != len(fields):
			raise Exception("bad nfields")  # todo: graceful
		opcount.append(fields[:-1])
		size.append(fields[-1])


X = numpy.asarray(opcount, dtype=numpy.float64)
Y = numpy.asarray(size, dtype=numpy.float64)

# print(X.shape, Y.shape)

nelem = X.shape[1]

m_value = numpy.asarray([1 for _ in range(nelem)], dtype=numpy.float64)

m = theano.shared(m_value, name='m')

x = T.matrix('x')
y = T.vector('y')

num_samples = X.shape[0]

prediction = T.dot(x, m.T)
cost = T.sum(T.pow(prediction-y, 2)) / (2*num_samples)

gradm = T.grad(cost, m)

learning_rate = 0.1
# training_steps = 100
training_steps = 3000

train = theano.function(
	[x, y],
	cost,
	updates=[(m, m-learning_rate*gradm)])
test = theano.function([x],prediction)

for i in range(training_steps):
    costM = train(X, Y)
    if i % 200 == 0:
	    print(costM)

print(costM)

mv = m.get_value()

for i in range(nfields-1):
	origv = mv[i]
	v = origv
	v = round(v)
	if v < 0:
		v = 0
	v = int(v)
	print("O"+ops[i], ":", v, ", //", origv)
