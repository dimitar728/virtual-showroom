export type Hotspot = {
  id: string;
  showroom_id: string;
  label: string;
  description?: string;
  pos_x: number;
  pos_y: number;
  pos_z: number;
  media_url?: string;
  link_url?: string;
  created_at?: string;
  updated_at?: string;
};

export type Booking = {
  id: string;
  showroom_id: string;
  user_id: string;
  start_time: string; // ISO string
  end_time: string;   // ISO string
  created_at?: string;
  updated_at?: string;
};