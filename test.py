import subprocess

commands = [
    ['./calcli', '-answer', '3'],
    ['./calcli', '-answer', '3+3'],
    ['./calcli', '-answer', '3^{2}'],
    ['./calcli', '-answer', '2+2+2+2+2+2'],
    ['./calcli', '-answer', '5^{1}+20-10'],
    ['./calcli', '-answer', '18*0+3^{3}-9'],
    ['./calcli', '-answer', '1*10+4^{2}/2+3'],
    ['./calcli', '-answer', '48/2+8^{0}-1'],
    ['./calcli', '-answer', '30-3'],
    ['./calcli', '-answer', '1^{0}+2^{2}+3^{2}+4^{2}']
]

for com in commands:
    print(subprocess.check_output(com))
