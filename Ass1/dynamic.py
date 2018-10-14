# make it work in both Python2 and 3
try:
    xrange
except NameError:
    xrange = range

def escapes(cr, ci, it):
    """
    Does iterating z <- z^2 + c escape after it iterations?
    """
    zr = 0.0
    zi = 0.0
    for i in xrange(it):
        # z <- z^2 + c
        zr,zi = zr*zr - zi*zi + cr, 2*zr*zi + ci
        if zr*zr + zi*zi > 4:
            return True
    return False

def toChar(p):
    if p:
        return " "
    else:
        return "X"

def mandel(xmin,xmax,xstep, ymin,ymax,ystep, iterations):
  for yc in xrange(ystep):
      y = yc*(ymax-ymin)/ystep + ymin
      row = []
      for xc in xrange(xstep):
          x = xc*(xmax-xmin)/xstep + xmin
          row.append( escapes(x, y, iterations) )
      
      print("".join([toChar(p) for p in row]))

