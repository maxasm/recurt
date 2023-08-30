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

const TextDialog = ({open, update_modal_open})=> {

    function handleModalClose(){
        update_modal_open(false) 
    }
    
    return (
        <Dialog
            open={open}
            onClose={handleModalClose}
            maxWidth={"xl"}
            scroll={"paper"}>
            <DialogTitle sx={{display: "flex", alignItems: "center"}}> 
                <Typography sx={{marginRight: "12px", fontSize: "1.2rem", fontFamily: "Barlow"}}> Rewriting block 1 of 10 </Typography>
                <CircularProgress sx={{marginBottom: "10px"}} size="25px"/>
            </DialogTitle>
            <DialogContent dividers>
                <DialogContentText>
                    <Typography>
                        Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros. Praesent commodo cursus magna, vel scelerisque nisl consectetur et. Cras mattis consectetur purus sit amet fermentum. Cras justo odio, dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta ac consectetur ac, vestibulum at eros.
                    </Typography>
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button variant="outlined" sx={{fontFamily: "Roboto"}}> Copy Text </Button>
                <Button variant="contained" sx={{fontFamily: "Roboto"}}> Rewrite Again </Button>
            </DialogActions>
        </Dialog>
    )
}

export default TextDialog;
