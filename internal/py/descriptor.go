package py

const Descriptor string = `
import shelve
def main():
    db = shelve.open("%s.shelf")
    dkeys = list(db.keys())
    dkeys.sort()
    for x in dkeys:
        print ( x , db[ x ])
    db.close()
    return

if __name__ == "__main__":
    main()
`
