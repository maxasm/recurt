import "../css/App.css"

/** mui components **/
import CssBaseline from "@mui/material/CssBaseline"
import Typography from "@mui/material/Typography"
import Paper from "@mui/material/Paper"
import TextField from "@mui/material/TextField"
import CircularProgress from "@mui/material/CircularProgress"
import Button from "@mui/material/Button"
import Box from "@mui/material/Box"
import Link from "@mui/material/Link"

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

/** count words **/
import {wordsCount} from "words-count"

/** WebSocket connection **/
let ws_connection;

const App = ()=> {
    // 'Hire Me' modal state
    const [hm_modal_open, update_hm_modal_open] = useState(false)
    
    // 'Text Rewrite' modal state 
    const [tx_modal_open, update_tx_modal_open] = useState(false)
    
    // the header text of the text modal
    const [tx_modal_header_text, update_tx_modal_header_text] = useState("Scanning for AI text")
    
    // text value in the 'text-area'
    const [text, updateText] = useState("")
     
    // state for the number of words
    const [text_count, update_text_count] = useState(0)
    
    // state for 'done' rewriting
    const [done, updateDone] = useState(false)
        
    // state for modal text
    const [modal_text, update_modal_text] = useState("")

    function handleOnTextChange(e) {
        updateText(e.target.value) 
        update_text_count(()=> {
            return wordsCount(e.target.value)
        })
    }
    
    // text value error
    const [text_error, update_text_error] = useState(false)
    const [text_error_msg, update_text_error_msg] = useState(false)

    function validateText() {
        const MAX_WORDS = 300;
        const MIN_WORDS = 30;
        
        const word_count = wordsCount(text) 
    
        // check if the text is empty
        if (word_count === 0) {
            update_text_error(true)     
            update_text_error_msg("Please provide text; this field cannot be left empty")
            return false
        }
    
        if  (word_count > MAX_WORDS) {
            update_text_error(true) 
            update_text_error_msg("Please limit your input to under 300 words")
            return false
        }
        
        if  (word_count < MIN_WORDS) {
            update_text_error(true) 
            update_text_error_msg("Please enter a minimum of 30 words")
            return false
        }
        
        update_text_error(false)
        update_text_error_msg("")
        
        update_modal_text(text)
        return true
    }
 
    // function called to rewrite the text
    function handleRewriteText(again) {
        // start by validating the text 
        let valid = validateText() 
        if (valid) {
            let ok = sendRewriteRequest()
            update_tx_modal_open(ok)
    
            if (again) {
                updateDone(false)
                update_tx_modal_header_text("Scanning for AI text")
            }
        }
    }

    // helper function for closing the WebSocket connection
    function closeWebSocketConnection() {
        if (ws_connection.readyState === "OPEN") {
            ws_connection.close() 
        }        
    } 
    

    async function sendRewriteRequest() {
        const resp = await fetch("/rewrite", {
            method: "POST",
            body: text, 
        }) 
    
       let resp_json = await resp.json() 
    
        if (resp_json.code !== 200) {
            console.log("There was an issue contacting the server")
            return false 
        }
    
        // get the unique id 
        let rw_id = resp_json.text
        
        // TODO: Put in the URL -> ai.maaax.pro 
        // create a new websocket connection
        ws_connection = new WebSocket(`ws://127.0.0.1:8080/ws/${rw_id}`)
        
        ws_connection.onopen = function() {
            console.log("websocket connected")     
        }
    
        ws_connection.onmessage = async function({data}) {
            let data_json = JSON.parse(data)
            // update the done state
            updateDone(data_json.done)
    
            if (data_json.done) {
                update_modal_text(data_json.text) 
                return
            }

            // update the modal header text
            update_tx_modal_header_text(data_json.text)
        }
    
        ws_connection.onclose = function() {
            console.log("websocket connection closed")
        } 
    
        ws_connection.onerror = function() {
            console.log(`websocket connection error: ${err}`)
        }
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
                <TextDialog
                    open={tx_modal_open}
                    update_modal_open={update_tx_modal_open}
                    base_text={modal_text}
                    header_text={tx_modal_header_text}
                    done={done}
                    rewriteFn={handleRewriteText} 
                    closeWebSocketConnection={closeWebSocketConnection}
                />

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
                    <Typography sx={{marginTop: "40px"}} variant="h1"> Bypass Turnitin AI Plagiarism detection in Seconds </Typography> 
                    <Typography variant="h2"> The best tool for Kenyan academic writers </Typography>
                    <Paper sx={{width: "90%", minHeight:"300px", margin: "0px auto", marginTop: "35px", paddingBottom: "10px", background: "rgba(255,255,255,0.5)", filter: "50px", display: "flex", flexDirection: "column", alignItems: "stretch", justifyContent:"center"}}>
                        <div className="instructions">
                            <Typography sx={{marginBottom: "10px", marginTop: "5px"}} variant="h3">📌 Tips on how to use the tool </Typography>
                            <Typography variant="ins">👉 Only rewrite ONE paragraph at a time. </Typography>
                            <Typography variant="ins">👉 For performance, do NOT include the title of the paragraph, just enter the text content. </Typography>
                            <Typography variant="ins">👉 To confirm that the generated text is indeed human text, you can use Turnitin or <a href="https://www.zerogpt.com/" target="_blank" sx={{cursor: "pointer"}}> ZeroGPT </a> </Typography>
                        </div>
                        <TextField 
                            variant="outlined"
                            fontFamily="Barlow"
                            multiline
                            rows={7}
                            value={text}
                            onChange={handleOnTextChange}
                            error={text_error}
                            helperText={text_error_msg} 
                            placeholder="Input your text. Minimum: 30 words. Maximum: 300 words."
                            label="Paste in your GPT paragraph"
                            sx={{width: "90%", margin: "15px auto"}}
                        /> 
                        <div className="btn-holder">
                            <Typography sx={{marginRight: "10px"}}> {`${text_count}/300`} </Typography>
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
