from PIL import Image



inp = []
with open("puzzle4input.txt") as mapfile:
    inp = mapfile.readlines()

y = len(inp)
x = len(inp[0])

xCol = (128, 0, 0)
mCol = (0, 128, 200)
aCol = (0, 0, 128)
sCol = (128, 128, 0)
matchCol = (0, 255, 0)
diagMatchCol = (0, 0, 255)

matchImg = Image.new("RGB", (8, 8), matchCol)
xImg = Image.open("puzzle4/x.png")
mImg = Image.open("puzzle4/m.png")
aImg = Image.open("puzzle4/a.png")
sImg = Image.open("puzzle4/s.png")

# matchX = Image.blend(xImg, Image.new("RGB", (8,8), matchCol), 0.5)
# matchM = Image.blend(mImg, Image.new("RGB", (8,8), matchCol), 0.5)
# matchA = Image.blend(aImg, Image.new("RGB", (8,8), matchCol), 0.5)
# matchS = Image.blend(sImg, Image.new("RGB", (8,8), matchCol), 0.5)


def directionMatch(direction, img):
    match direction:
        case "NORTH": return Image.blend(img, Image.new("RGB", (8, 8), matchCol), 0.5)
        case "SOUTH": return Image.blend(img, Image.new("RGB", (8, 8), matchCol), 0.5)
        case "EAST": return Image.blend(img, Image.new("RGB", (8, 8), matchCol), 0.5)
        case "WEST": return Image.blend(img, Image.new("RGB", (8, 8), matchCol), 0.5)
        case "NORTH_WEST": return Image.blend(img, Image.new("RGB", (8, 8), diagMatchCol), 0.5)
        case "NORTH_EAST": return Image.blend(img, Image.new("RGB", (8, 8), diagMatchCol), 0.5)
        case "SOUTH_WEST": return Image.blend(img, Image.new("RGB", (8, 8), diagMatchCol), 0.5)
        case "SOUTH_EAST": return Image.blend(img, Image.new("RGB", (8, 8), diagMatchCol), 0.5)
        


img = Image.new("RGB", ((x+1) * 8, (y+1) * 8), (255, 255, 255))

for y, row in enumerate(inp):
    for x, col in enumerate(row):
        match col:
            case "X": img.paste(xImg, (x * 8, y * 8))#img.putpixel((x, y), xCol)
            case "M": img.paste(mImg, (x * 8, y * 8))#img.putpixel((x, y), mCol)
            case "A": img.paste(aImg, (x * 8, y * 8))#img.putpixel((x, y), aCol)
            case "S": img.paste(sImg, (x * 8, y * 8))#img.putpixel((x, y), sCol)

pairs = []
points = []
with open("p4output.txt") as pairsFile:
    pairs = pairsFile.readlines()

for p in pairs:
    mat = p.split("|")
    pts = mat[1].split(";")
    for point in pts:
        ptx = point.split(",")
        points.append((mat[0], int(ptx[0]), int(ptx[1])))

#print(points)

for point in points:
    match inp[point[2]][point[1]]:
        case "X": img.paste(directionMatch(point[0], xImg), (point[1] * 8, point[2] * 8)) 
        case "M": img.paste(directionMatch(point[0], mImg), (point[1] * 8, point[2] * 8)) 
        case "A": img.paste(directionMatch(point[0], aImg), (point[1] * 8, point[2] * 8)) 
        case "S": img.paste(directionMatch(point[0], sImg), (point[1] * 8, point[2] * 8)) 
    
    #img.putpixel(point, matchCol)

img.save("puzzle4map.png")