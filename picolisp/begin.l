(load "@lib/math.l")

(de Begin1 (A)
   (* 4 A))

(de Begin2 (A)
   (* A A))

(de Begin3 (A B)
   (list (* A B) (* 2 (+ A B))))

(de Begin4 (D)
   (* D 3.14))

(de Begin5 (A)
   (let A2 (* A A)
      (list (* A A2) (* A 6))))

(de Begin6 (A B C)
   (list (* A B C)
      (* 2 (+ (* A B) (* B C) (* A C)))))

## (de
