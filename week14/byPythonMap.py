counts  =dict()
print(counts)


for i in range(5):
    name = input("후보자 이름 입력:")
    counts[name] = i+1


print(counts)
# ===================================================

# lines = []
# with open("vote.txt", "r") as fp:
#     lines.append(fp.readline)

# names = list()
# counts = []

# for line in lines:
#     matched = False
#     for name in names:
#         if name == line:
#             counts[i] = counts[i] + 1
#             matched = True
#         i =  i+1
#     if matched == False:
#         names.append(line)
#         counts.append(1)
# for name in names:
#     print(f"{name}: {counts[i]}")
#     i = i+1


lines = []
with open("votes.txt", "r") as fp:
    lines = fp.readlines()

names = []
counts = []

for line in lines:
    matched = False
    for i, name in enumerate(names):
        if name.strip() == line.strip():
            counts[i] = counts[i] + 1
            matched = True
            break
    if not matched:
        names.append(line.strip())
        counts.append(1)

for i, name in enumerate(names):
    print(f"{name}: {counts[i]}")
