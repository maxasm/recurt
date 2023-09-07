/** mui components **/
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Typography from "@mui/material/Typography";
import DialogActions from '@mui/material/DialogActions';
import Button from "@mui/material/Button"
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';

/** icons **/
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

/** state **/
import {useState} from "react"

/** colors **/
import {grey} from "@mui/material/colors"

const HireMeDialog = ({open, update_modal_open})=> {

    function handleOnModalClose() {
        update_modal_open(false)
    }
    
    // TODO: use an <Accordion/> for the job descriptions 
    return (
        <Dialog
            open={open}
            PaperProps={{sx:{background: grey[900], color: "#ffffff"}}}
            onClose={handleOnModalClose}>

            <DialogTitle>
                <Typography> Hi there 👋, I am Maxwell - a Kenyan Sofware Dev 🛠️ </Typography>
            </DialogTitle>
    
            <DialogContent dividers>
                <DialogContentText sx={{color: "#ffffff"}}>
                    <Accordion sx={{background: grey[800], color: "#ffffff"}}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon sx={{color: "#ffffff"}}/>}
                        >
                            <Typography>
                                ⚙️  Get a new custom web site
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            Get a new custom website 💻 for your brand or business for as low as 25,000 KSH. Price is inclusive of UI/UX design, Front End, Backend, Hosting fees and a custom domain*
                        </AccordionDetails>
                    </Accordion>
                    <Accordion sx={{background: grey[800], color: "#ffffff"}}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon sx={{color: "#ffffff"}}/>}
                        >
                            <Typography>
                                ⚙️  Get a new Android App 
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            Get a new custom Android application 🎮 for your brand or business from as low as 20,000 KSH. 
                        </AccordionDetails>
                    </Accordion>
                    <Accordion sx={{background: grey[800], color: "#ffffff"}}>
                        <AccordionSummary
                            expandIcon={<ExpandMoreIcon sx={{color: "#ffffff"}}/>}
                        >
                            <Typography>
                                ⚙️  Technical writing 
                            </Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                            I am also an Academic writer 😉 
                        </AccordionDetails>
                    </Accordion>

                </DialogContentText>
            </DialogContent>
    
            <DialogActions>
                <Typography>
                    👉 awsmaaax@gmail.com (email me)
                </Typography> 
            </DialogActions>
        </Dialog>
    );

}

export default HireMeDialog;
