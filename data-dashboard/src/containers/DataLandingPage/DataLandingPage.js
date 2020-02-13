import React, {Component} from "react";
import {withStyles} from "@material-ui/core";
import PropTypes from "prop-types";
import Paper from "@material-ui/core/Paper/Paper";
import Grid from "@material-ui/core/Grid/Grid";
import Typography from "@material-ui/core/Typography/Typography";
import Button from '@material-ui/core/Button';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';


const styles = theme => ({
    body: {
        maxWidth: theme.spacing.getMaxWidth.maxWidth,
        margin: theme.spacing.getMaxWidth.margin,
    },
    paper: {
        padding: theme.spacing.unit * 2,
        textAlign: 'left',
        color: theme.palette.text.secondary,
    },
    paperImage: {
        textAlign: 'center',
        padding: theme.spacing.unit * 2,
    },
    metImage: {
        maxWidth: '100%',
    },
    '& > *': {
        margin: theme.spacing.unit * 5,
    },
});

class DataLandingPage extends Component {

    render() {
        const { classes } = this.props;
        return (
            <div className={classes.body}>
                <Grid container spacing={24}>
                    <Grid item xs={12}>
                        <Paper className={classes.paper}>
                            <Typography variant="headline" gutterBottom>
                                Arome Arctic
                            </Typography>
                            <Typography variant="h6" gutterBottom>
                                Description
                            </Typography>
                            <Typography variant="body1" paragraph>
                                Description, description, Description, description, Description, description, Description, description, Description, description
                            </Typography>
                            <Typography variant="body1" paragraph>
                                Some more description, description
                            </Typography>
                            <Typography variant="h6" gutterBottom>
                                Metadata
                            </Typography>
                            <Typography variant="body1" paragraph>
                                <List>
                                    <ListItem>
                                        <ListItemText primary="Bounding box" secondary="10, 90, 10, 90"/>
                                    </ListItem>
                                    <ListItem>
                                        <ListItemText primary="Last Metadata Update" secondary="2019-12-12T12:30:01Z"/>
                                    </ListItem>
                                    <ListItem>
                                        <ListItemText primary="Temporal Extent" secondary="Dataset series started on 2019-09-12T12:30:01Z and are updated 4 times pr day."/>
                                    </ListItem>
                                </List>
                            </Typography>
                            <Typography variant="h6" gutterBottom>
                                Code examples
                            </Typography>
                            <Typography variant="body1" paragraph>
                                Code examples for accessing data for this dataset series. 
                            </Typography>
                            <Button component="span">Python</Button>
                            <Button component="span">R</Button>
                            <Button component="span">Go</Button>
                            <Typography variant="h6" gutterBottom>
                                Access services
                            </Typography>
                            <Button component="span">WMS</Button>
                            <Button component="span">OPeNDAP 4</Button>
                            <Button component="span">HTTP Download</Button>
                        </Paper>
                    </Grid>
                </Grid>
            </div>
        );
    }

}

DataLandingPage.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DataLandingPage);
