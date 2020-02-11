import React, {Component} from "react";
import {withStyles, TextField} from "@material-ui/core";
import PropTypes from "prop-types";
import Paper from "@material-ui/core/Paper/Paper";
import Grid from "@material-ui/core/Grid/Grid";
import Typography from "@material-ui/core/Typography/Typography";
import {FormControl} from '@material-ui/core';
import Button from '@material-ui/core/Button';


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
    }
});

class Register extends Component {

    render() {
        const { classes } = this.props;
        return (
            <div className={classes.body}>
                <Grid container spacing={24}>
                    <Grid item xs={12}>
                        <Paper className={classes.paper}>
                            <Typography variant="headline" gutterBottom>
                                Register a new dataset
                            </Typography>
                            <FormControl>
                                <TextField 
                                    id="title"
                                    label="Title"
                                    helperText="Follow this format: blah"
                                />
                                 <TextField
                                    id="abstract"
                                    label="Abstract"
                                    helperText="Short and concise description of dataset."
                                />
                                <TextField
                                    id="datasetURI"
                                    label="datasetURI"
                                    helperText="OpenDAP or local file URI to a representative dataset."
                                />
                            </FormControl>
                        </Paper>
                    </Grid>
                    <Grid item xs={3} justify="center">
                        <Paper className={classes.paper}>
                            <label htmlFor="text-button-file">
                                <Button component="span">Register dataset</Button>
                            </label>
                        </Paper>
                    </Grid>
                </Grid>
            </div>
        );
    }

}

Register.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Register);
