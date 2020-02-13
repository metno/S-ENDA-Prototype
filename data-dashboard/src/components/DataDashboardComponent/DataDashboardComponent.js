import React, {Component} from "react";
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import homePageImg from './homePageImg.png';
import Typography from "@material-ui/core/Typography/Typography";
import Link from '@material-ui/core/Link';


const styles = theme => ({
    root: {
        flexGrow: 1,
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
    homePageImg: {
        maxWidth: '100%',
    }
});


class DataDashboardComponent extends Component {

    state = {
        persons: []
    };

    render() {
        const { classes } = this.props;
        return (
            <div className={classes.root}>
                <Grid container spacing={24}>
                    <Grid item xs={12}>
                        <Paper className={classes.paperImage}>
                            <Typography variant="headline" gutterBottom>
                                Your Data Dashboard
                            </Typography>
                        </Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.paper}>
                            <Typography variant="headline" gutterBottom>
                            <Link href="/datalandingpage">Arome Arctic</Link>
                            </Typography>
                            <Typography>
                                <p>
                                    <strong>Operations</strong>: On time and correct in Dataroom A and Dataroom B.
                                </p>
                                <p>
                                    <strong>Usage</strong>: 1001 DAP requests daily, 502 map requests.
                                </p>
                            </Typography>
                        </Paper>
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <Paper className={classes.paper}>
                            <Typography variant="headline" gutterBottom>
                                Arome Arctic experimental uber version    
                            </Typography>
                            <Typography>       
                                <p>
                                    <strong>Operations</strong>: Last run correct, but late by 2 hours in Dataroom B.
                                </p>
                                <p>
                                    <strong>Usage</strong>: 2 DAP requests daily, 0 map requests.
                                </p>
                            </Typography>
                        </Paper>
                    </Grid>
                    <Grid item xs={6} sm={3}>
                        <Paper className={classes.paper}>
                            <Typography variant="headline" gutterBottom>
                                <Link  href="/dataset">Register new dataset series</Link>
                            </Typography>
                        </Paper>
                    </Grid>
                    <Grid item xs={6} sm={3}>
                        <Paper className={classes.paper}>
                            <Typography variant="headline" gutterBottom>
                                <Link href="#metrics">Statistics about data production and usage</Link>
                            </Typography>
                        </Paper>
                    </Grid>
                </Grid>
            </div>
        );
    }
}

DataDashboardComponent.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(DataDashboardComponent);