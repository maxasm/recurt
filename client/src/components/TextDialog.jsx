/** mui components **/
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Dialog from "@mui/material/Dialog"
import Button from "@mui/material/Button"
import Link from "@mui/material/Link"
import Typography from "@mui/material/Typography"
import CircularProgress from '@mui/material/CircularProgress'

/** state management **/
import {useState} from "react"

// TODO: try and use a <pre/> tag to maintain spaces.
const TextDialog = ({open, update_modal_open, base_text, done, header_text})=> {

    function handleModalClose(){
        update_modal_open(false) 
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
                <Button disabled={!done} variant="outlined" sx={{fontFamily: "Roboto"}}> Copy Text </Button>
                <Button disabled={!done} variant="contained" sx={{fontFamily: "Roboto"}}> Rewrite Again </Button>
            </DialogActions>
        </Dialog>
    )
}

export default TextDialog;
