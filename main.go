package main

//import "os"
// import "log"

func main() {
    queries := []string {
        "https://goldenhome.gr/property/index?PropertySearch%5BPropertyID%5D=&PropertySearch%5BTrnTypeID%5D=2&PropertySearch%5Bvideo_url%5D=&PropertySearch%5BPropCategID%5D=11704&category=&PropertySearch%5BPropSubCategID%5D=&PropertySearch%5BareaLevel1%5D=101000000&PropertySearch%5BRAreaID%5D=&PropertySearch%5BFloorNo%5D=&PropertySearch%5BFloorNo_to%5D=&PropertySearch%5BBuiltYear%5D=&PropertySearch%5BBuiltYear_to%5D=&PropertySearch%5BTotalRooms%5D=2&PropertySearch%5BTotalRooms_to%5D=5&PropertySearch%5BTotalParkings%5D=&PropertySearch%5BTotalParkings_to%5D=2&PropertySearch%5BAskedValue%5D=&PropertySearch%5BAskedValue_to%5D=1000000&PropertySearch%5BTotalSm%5D=80&PropertySearch%5BTotalSm_to%5D=&PropertySearch%5Bapothiki%5D=&PropertySearch%5Btzaki%5D=", // ΚΕΝΤΡΟ
        "https://goldenhome.gr/property/index?PropertySearch%5BPropertyID%5D=&PropertySearch%5BTrnTypeID%5D=2&PropertySearch%5Bvideo_url%5D=&PropertySearch%5BPropCategID%5D=11704&PropertySearch%5BPropSubCategID%5D=&PropertySearch%5BareaLevel1%5D=102000000&PropertySearch%5BRAreaID%5D=&PropertySearch%5BFloorNo%5D=&PropertySearch%5BFloorNo_to%5D=&PropertySearch%5BBuiltYear%5D=&PropertySearch%5BBuiltYear_to%5D=&PropertySearch%5BTotalRooms%5D=2&PropertySearch%5BTotalRooms_to%5D=5&PropertySearch%5BTotalParkings%5D=&PropertySearch%5BTotalParkings_to%5D=2&PropertySearch%5BAskedValue%5D=&PropertySearch%5BAskedValue_to%5D=1000000&PropertySearch%5BTotalSm%5D=80&PropertySearch%5BTotalSm_to%5D=&PropertySearch%5Bapothiki%5D=&PropertySearch%5Btzaki%5D=&category=", // ΒΟΡΕΙΑ ΠΡΟΑΣΤΙΑ
        "https://goldenhome.gr/property/index?PropertySearch%5BPropertyID%5D=&PropertySearch%5BTrnTypeID%5D=2&PropertySearch%5Bvideo_url%5D=&PropertySearch%5BPropCategID%5D=11704&category=&PropertySearch%5BPropSubCategID%5D=&PropertySearch%5BareaLevel1%5D=104000000&PropertySearch%5BRAreaID%5D=&PropertySearch%5BFloorNo%5D=&PropertySearch%5BFloorNo_to%5D=&PropertySearch%5BBuiltYear%5D=&PropertySearch%5BBuiltYear_to%5D=&PropertySearch%5BTotalRooms%5D=2&PropertySearch%5BTotalRooms_to%5D=5&PropertySearch%5BTotalParkings%5D=&PropertySearch%5BTotalParkings_to%5D=2&PropertySearch%5BAskedValue%5D=&PropertySearch%5BAskedValue_to%5D=1000000&PropertySearch%5BTotalSm%5D=80&PropertySearch%5BTotalSm_to%5D=&PropertySearch%5Bapothiki%5D=&PropertySearch%5Btzaki%5D=", // ΑΝΑΤΟΛΙΚΑ
    }
    getAds(queries)
}
