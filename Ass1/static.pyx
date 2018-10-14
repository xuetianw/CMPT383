def escapes(double cr, double ci, int it):
    """
    Does iterating z <- z^2 + c escape after it iterations?
    """
    cdef double zr = 0.0
    cdef double zi = 0.0
    cdef int i
    for i in range(it):
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
  for yc in range(ystep):
      y = yc*(ymax-ymin)/ystep + ymin
      row = []
      for xc in range(xstep):
          x = xc*(xmax-xmin)/xstep + xmin
          row.append( escapes(x, y, iterations) )
      
      print("".join([toChar(p) for p in row]))

