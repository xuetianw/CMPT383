def count_occurrences(lst, val):
    count = 0
    for v in lst:
        if v == val:
            count += 1
    return count




def main():
	thislist1 = [1,2,3]
	thislist2 = {1,2,3}
	a = count_occurrences (thislist1, 0)
	b = count_occurrences (thislist2, 0)
	print (type(thislist1))
	print (type(thislist2))
	print (a)
	print (b)
  
if __name__== "__main__":
  main()

