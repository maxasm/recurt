import "../css/App.css"

/** mui components **/
import CssBaseline from "@mui/material/CssBaseline"
import Typography from "@mui/material/Typography"
import Paper from "@mui/material/Paper"
import TextField from "@mui/material/TextField"
import CircularProgress from "@mui/material/CircularProgress"
import Button from "@mui/material/Button"
import Box from "@mui/material/Box"

/** components **/
import TextDialog from "./TextDialog"
import HireMeDialog from "./HireMeDialog"

/** colors **/
import {blue, grey} from "@mui/material/colors"

/** icons **/
import ModelTrainingIcon from '@mui/icons-material/ModelTraining';
import AddIcCallIcon from '@mui/icons-material/AddIcCall';

/** themes **/
import {createTheme, ThemeProvider} from "@mui/material/styles"

/** state **/
import {useState} from "react"

const App = ()=> {
    // 'Hire Me' modal state
    const [hm_modal_open, update_hm_modal_open] = useState(false)
    
    // 'Text Rewrite' modal state 
    const [tx_modal_open, update_tx_modal_open] = useState(false)
    
    // function called to rewrite the text
    function handleRewriteText() {
        // open the text modal
        // TODO: start the rewriting procedure        
        update_tx_modal_open(true)
    }

    // open the 'hire me' dialog
    function handleOpenHireMeDialog() {
        update_hm_modal_open(true)
    }  

    const theme = createTheme({
        typography: {
            fontFamily: "Overpass",
            color: grey[900],
            h1: {
                fontSize: "2.5rem",
                fontWeight: "600",
                textAlign: "center",
            },
            h2: {
                fontSize: "1.4rem", 
                fontWeight: "400",
                textAlign: "center",
                fontFamily: "Barlow",
            },
            h3: {
                fontSize: "1.2rem", 
                fontWeight: 600,
            },
            ins: {
                fontSize: "0.8rem",
            }, 
        },
    });

    return (
        <>
            <ThemeProvider theme={theme}>
                <CssBaseline/>
                <TextDialog open={tx_modal_open} update_modal_open={update_tx_modal_open}/>
                <HireMeDialog open={hm_modal_open} update_modal_open={update_hm_modal_open}/>

                <div className="nav">
                    <Typography sx={{fontSize: "1.4rem", fontFamily: "logo"}}> Maaax.pro </Typography>
                    <Button
                        onClick={handleOpenHireMeDialog}
                        startIcon={<AddIcCallIcon/>}
                        variant="outlined"
                        size="small"
                        sx={{fontFamily: "Barlow", fontWeight: "500"}}> Hire me as a dev </Button>
                </div>
                <div className="container">
                    <Typography sx={{marginTop: "40px"}} variant="h1"> Bypass Turnitin AI and Plagiarism detection in minutes </Typography> 
                    <Typography variant="h2"> The best tool for Kenyan academic writers </Typography>
                    <Paper sx={{width: "90%", minHeight:"300px", margin: "0px auto", marginTop: "35px", paddingBottom: "10px", background: "rgba(255,255,255,0.5)", filter: "50px", display: "flex", flexDirection: "column", alignItems: "stretch", justifyContent:"center"}}>
                        <div className="instructions">
                            <Typography sx={{marginBottom: "10px", marginTop: "5px"}} variant="h3">📌 Tips on how to use the tool </Typography>
                            <Typography variant="ins">👉 Only rewrite ONE paragraph at a time. </Typography>
                            <Typography variant="ins">👉 For performance, do NOT include the title of the paragraph, just enter the text content. </Typography>
                            <Typography variant="ins">👉 The tool thoroughly reviews and rewrites your text to eliminate any AI-generated content. Please be patient. </Typography>
                        </div>
                        <TextField 
                            variant="outlined"
                            fontFamily="Barlow"
                            multiline
                            rows={7}
                            placeholder="Input your text. Minimum: 200 words. Maximum: 300 words."
                            label="Paste in your GPT paragraph"
                            sx={{width: "90%", margin: "15px auto"}}
                        /> 
                        <div className="btn-holder">
                            <Button
                                onClick={handleRewriteText}
                                startIcon={<ModelTrainingIcon/>}
                                size="medium"
                                sx={{fontFamily: "Roboto", fontWeight: "500"}}
                                variant="contained"> Rewrite Text </Button>
                        </div>
                    </Paper>
                </div>
            </ThemeProvider>
        </>
    );
}

export default App;
