func testSelectFor2(chExit chan bool){
	EXIT:
	   for  {
		   select {
		   case v, ok := <-chExit:
			   if !ok {
				   fmt.Println("close channel 2", v)
				   break EXIT//goto EXIT2
			   }
   
			   fmt.Println("ch2 val =", v)
		   }
	   }
   
	   //EXIT2:
	   fmt.Println("exit testSelectFor2")
   }

// 通常在for循环中，使用break可以跳出循环，但是注意在go语言中，for select配合时，break 并不能跳出循环。