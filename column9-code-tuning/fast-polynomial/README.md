#### 9.5.12

To evaluate the polynomial
y = a_nx^n + a_n–1x^(n –1) + ... + a_1x^1 + a_0

the following code uses 2n multiplications. Give a faster function.</br>
y = a[0]</br>
xi = 1</br>
for i = [1, n]</br>
    xi = x * xi</br>
    y = y + a[i]*xi

