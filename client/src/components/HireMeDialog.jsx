/** mui components **/
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Typography from "@mui/material/Typography";
import DialogActions from '@mui/material/DialogActions';
import Button from "@mui/material/Button"

/** state **/
import {useState} from "react"

/** colors **/
import {grey} from "@mui/material/colors"

const HireMeDialog = ({open, update_modal_open})=> {

    function handleOnModalClose() {
        update_modal_open(false)
    }
    
    return (
        <Dialog
            open={open}
            PaperProps={{sx:{background: grey[900], color: "#ffffff"}}}
            onClose={handleOnModalClose}>

            <DialogTitle>
                <Typography> Hi there 👋, I am Maxwell - a Sofware Dev 🛠️ </Typography>
            </DialogTitle>
    
            <DialogContent dividers>
                <DialogContentText sx={{color: "#ffffff"}}>
                    <Typography>
                        ⚙️  Get a new custom web site (+Domain and Hosting) from as low as 25,000 KSH
                    </Typography>
    
                    <Typography>
                        ⚙️  Get a new custom Android App from as low as 20,000 KSH
                    </Typography>
                </DialogContentText>
            </DialogContent>
    
            <DialogActions>
                <Typography>
                     📡 maxdev@maaax.pro (email me)
                </Typography> 
            </DialogActions>
        </Dialog>
    );

}

export default HireMeDialog;
