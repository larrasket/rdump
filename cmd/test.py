import shelve

def main():
    db = shelve.open("Mahdi_Amel.shelf")
    dkeys = list(db.keys())
    dkeys.sort()
    for x in dkeys:
        print ( x , db[ x ])
    db.close()
    return

if __name__ == "__main__":
    main()
