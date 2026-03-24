import * as React from 'react';
import { Class, Email, Grade, Person, Phone, Download } from '@mui/icons-material';
import { Box, Card, CardContent, Divider, Grid2, Typography, Button } from '@mui/material';
import { useDownloadStudentReportMutation } from '../../api/student-api';

type MiniAvatarProps = {
  id: number | string;
  name: string;
  phone: string;
  email: string;
  selectedClass: string;
  section: string;
};

export const MiniAvatar: React.FC<MiniAvatarProps> = ({
  id,
  name,
  selectedClass,
  section,
  phone,
  email
}) => {
  const [downloadReport, { isLoading }] = useDownloadStudentReportMutation();

  const handleDownloadReport = async () => {
    try {
      const blob = await downloadReport(id.toString()).unwrap();
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.download = `student_${id}_report.pdf`;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error('Error downloading report:', error);
      alert('Failed to download report. Please try again.');
    }
  };
  return (
    <Card variant='outlined'>
      <CardContent>
        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
          <Person sx={{ mr: 1 }} />
          <Typography variant='h6'>{name}</Typography>
        </Box>
        <Divider sx={{ mb: 2 }} />
        <Grid2 container spacing={2}>
          <Grid2 size={{ xs: 12 }}>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Class sx={{ mr: 1 }} />
              <Typography variant='subtitle2'>Class</Typography>
            </Box>
            <Typography variant='body1'>{selectedClass}</Typography>
          </Grid2>
          <Grid2 size={{ xs: 12 }}>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Grade sx={{ mr: 1 }} />
              <Typography variant='subtitle2'>Section</Typography>
            </Box>
            <Typography variant='body1'>{section}</Typography>
          </Grid2>
          <Grid2 size={{ xs: 12 }}>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Email sx={{ mr: 1 }} />
              <Typography variant='subtitle2'>Email</Typography>
            </Box>
            <Typography variant='body1'>{email}</Typography>
          </Grid2>
          <Grid2 size={{ xs: 12 }}>
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <Phone sx={{ mr: 1 }} />
              <Typography variant='subtitle2'>Phone</Typography>
            </Box>
            <Typography variant='body1'>{phone}</Typography>
          </Grid2>
          <Grid2 size={{ xs: 12 }}>
            <Button
              variant='contained'
              startIcon={<Download />}
              onClick={handleDownloadReport}
              disabled={isLoading}
              fullWidth
              sx={{ mt: 2 }}
            >
              {isLoading ? 'Generating Report...' : 'Download Report'}
            </Button>
          </Grid2>
        </Grid2>
      </CardContent>
    </Card>
  );
};
