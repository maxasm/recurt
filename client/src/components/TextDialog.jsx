/** mui components **/
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Dialog from "@mui/material/Dialog";
import Button from "@mui/material/Button";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";
import CircularProgress from '@mui/material/CircularProgress';
import Snackbar from '@mui/material/Snackbar';

/** copy text **/
import copy from "clipboard-copy"

/** state management **/
import {useState} from "react"

// TODO: try and use a <pre/> tag to maintain spaces.
const TextDialog = ({open, update_modal_open, base_text, done, header_text, rewriteFn})=> {

    // 'snackbar' open state
    const [sb_open, update__sb_open] = useState(false)
    
    // snackbar message
    const [sb_message, update__sb_message] = useState("")

    // handle snackbar closing 
    function handleOnSbClose() {
        update__sb_open(false) 
    }

    function handleModalClose(){
        update_modal_open(false) 
    }
        
    function handleCopyText() {
        // use 'dom' api
        copy(base_text)
        update__sb_message("Text copied successfully.")
        update__sb_open(true)
    }
    
    function handleRewriteAgain() {
        update__sb_message("Rewriting text again ... ")
        update__sb_open(true)
        rewriteFn(true) 
    }
    
    // TODO: remove <Typography/> inside <DialogTitle/>
    return (
        <Dialog
            open={open}
            onClose={handleModalClose}
            maxWidth={"xl"}
            scroll={"paper"}>
            <DialogTitle sx={{display: "flex", alignItems: "center"}}> 
                <Typography
                    sx={{marginRight: "12px", fontSize: "1.2rem", fontFamily: "Barlow"}}> {!done ? header_text : "Human Rewritten Text"} </Typography>
                <CircularProgress
                    sx={{display: !done ? "inline" : "none"}}
                    size="25px"/>
            </DialogTitle>
            <DialogContent dividers>
                <DialogContentText>
                    <Typography>
                        {base_text}
                    </Typography>
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button
                    onClick={handleCopyText}
                    disabled={!done}
                    variant="outlined"
                    sx={{fontFamily: "Roboto"}}> Copy Text </Button>
                <Button
                    onClick={handleRewriteAgain}
                    disabled={!done}
                    variant="contained"
                    sx={{fontFamily: "Roboto"}}> Rewrite Again </Button>
            </DialogActions>
            <Snackbar
                open={sb_open} 
                message={sb_message}
                autoHideDuration={3000}
                onClose={handleOnSbClose}
            />
        </Dialog>
    )
}

export default TextDialog;
