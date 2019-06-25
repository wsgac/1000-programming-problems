(load "@lib/math.l")

(de Begin1 (a)
   (* 4 a))

(de Begin2 (a)
   (* a a))

(de Begin3 (a b)
   (list (* a b) (* 2 (+ a b]

(de Begin4 (d)
   (* d 3.14]

(de Begin5 (a)
   (let a2 (* a a) # Save unnecessary multiplications
      (list (* a a2) (* a 6]

(de Begin6 (a b c)
   (list (* a b c)
         (* 2 (+ (* a b) (* b c) (* a c]


